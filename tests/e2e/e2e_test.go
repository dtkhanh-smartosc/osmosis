//go:build e2e
// +build e2e

package e2e

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"

	appparams "github.com/osmosis-labs/osmosis/v11/app/params"
	"github.com/osmosis-labs/osmosis/v11/tests/e2e/configurer/config"
	"github.com/osmosis-labs/osmosis/v11/tests/e2e/initialization"
)

func (s *IntegrationTestSuite) TestCreatePoolPostUpgrade() {
	if s.skipUpgrade {
		s.T().Skip("pool creation tests are broken when upgrade is skipped. To be fixed in #1843")
	}
	chain := s.configurer.GetChainConfig(0)
	node, err := chain.GetDefaultNode()
	s.NoError(err)

	node.CreatePool("pool2A.json", initialization.ValidatorWalletName)
	node.CreatePool("pool2B.json", initialization.ValidatorWalletName)
}

func (s *IntegrationTestSuite) TestIBCTokenTransfer() {
	if s.skipIBC {
		s.T().Skip("Skipping IBC tests")
	}

	chainA := s.configurer.GetChainConfig(0)
	chainB := s.configurer.GetChainConfig(1)

	chainA.SendIBC(chainB, chainB.NodeConfigs[0].PublicAddress, initialization.OsmoToken)
	chainB.SendIBC(chainA, chainA.NodeConfigs[0].PublicAddress, initialization.OsmoToken)
	chainA.SendIBC(chainB, chainB.NodeConfigs[0].PublicAddress, initialization.StakeToken)
	chainB.SendIBC(chainA, chainA.NodeConfigs[0].PublicAddress, initialization.StakeToken)
}

func (s *IntegrationTestSuite) TestIBCTokenTransferRateLimiting() {
	if s.skipIBC {
		s.T().Skip("Skipping IBC tests")
	}
	chainA := s.configurer.GetChainConfig(0)
	chainB := s.configurer.GetChainConfig(1)

	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	balance, _ := node.QueryBalances(chainA.NodeConfigs[0].PublicAddress)
	s.T().Log("CHAIN-A balance:", balance) // 96_700_000_000 uosmo

	// Sending >1%
	chainA.SendIBC(chainB, chainB.NodeConfigs[0].PublicAddress, sdk.NewInt64Coin(initialization.OsmoDenom, initialization.OsmoBalanceA*0.011))

	node.StoreWasmCode("rate_limiter.wasm", initialization.ValidatorWalletName)
	chainA.LatestCodeId += 1
	node.InstantiateWasmContract(strconv.Itoa(chainA.LatestCodeId), fmt.Sprintf("{\"gov_module\": \"%s\", \"ibc_module\": \"osmo1g7ajkk295vactngp74shkfrprvjrdwn662dg26\", \"paths\": [{\"channel_id\": \"channel-0\", \"denom\": \"%s\", \"quotas\": [{\"name\":\"testQuota\", \"duration\": 86400, \"send_recv\": [1, 1]}] } ] }", chainA.NodeConfigs[0].PublicAddress, initialization.OsmoToken.Denom), initialization.ValidatorWalletName)

	// Using code_id 1 because this is the only contract right now. This may need to change if more contracts are added
	contracts, err := node.QueryContractsFromId(chainA.LatestCodeId)
	s.T().Log(contracts)

	node.SubmitParamChangeProposal(fmt.Sprintf(`{"title":"Param change","description":"Changing rate limit contract param",
"changes":[{"subspace":"rate-limited-ibc","key":"contract","value":{"contract_address":"%s"}}],
"deposit":"10000000stake"}`, contracts[0]), initialization.ValidatorWalletName)
	chainA.LatestProposalNumber += 1

	for _, n := range chainA.NodeConfigs {
		n.VoteYesProposal(initialization.ValidatorWalletName, chainA.LatestProposalNumber)
	}

	s.Eventually(
		func() bool {
			noTotal, yesTotal, noWithVetoTotal, abstainTotal, err := node.QueryPropTally(chainA.LatestProposalNumber)
			if err != nil {
				return false
			}
			if abstainTotal.Int64()+noTotal.Int64()+noWithVetoTotal.Int64()+yesTotal.Int64() <= 0 {
				return false
			}
			return true
		},
		1*time.Minute,
		10*time.Millisecond,
		"Osmosis node failed to retrieve prop tally",
	)

	// Sending <1%. Should work
	chainA.SendIBC(chainB, chainB.NodeConfigs[0].PublicAddress, sdk.NewInt64Coin(initialization.OsmoDenom, initialization.OsmoBalanceA*0.08))

	// Sending >1%. Should fail
	chainA.FailIBC(chainB, chainB.NodeConfigs[0].PublicAddress, sdk.NewInt64Coin(initialization.OsmoDenom, initialization.OsmoBalanceA*0.011))

}

func (s *IntegrationTestSuite) TestSuperfluidVoting() {
	if s.skipUpgrade {
		// TODO: https://github.com/osmosis-labs/osmosis/issues/1843
		s.T().Skip("Superfluid tests are broken when upgrade is skipped. To be fixed in #1843")
	}
	const walletName = "superfluid-wallet"

	chain := s.configurer.GetChainConfig(0)
	node, err := chain.GetDefaultNode()
	s.NoError(err)

	// enable superfluid via proposal.
	node.SubmitSuperfluidProposal("gamm/pool/1", sdk.NewCoin(appparams.BaseCoinUnit, sdk.NewInt(config.InitialMinDeposit)))
	chain.LatestProposalNumber += 1
	node.DepositProposal(chain.LatestProposalNumber, false)
	for _, node := range chain.NodeConfigs {
		node.VoteYesProposal(initialization.ValidatorWalletName, chain.LatestProposalNumber)
	}

	walletAddr := node.CreateWallet(walletName)
	// send gamm tokens to node's other wallet (non self-delegation wallet)
	node.BankSend("100000000000000000000gamm/pool/1", chain.NodeConfigs[0].PublicAddress, walletAddr)
	// lock tokens from node 0 on chain A
	node.LockTokens("100000000000000000000gamm/pool/1", "240s", walletName)
	chain.LatestLockNumber += 1
	// superfluid delegate from non self-delegation wallet to validator 1 on chain.
	node.SuperfluidDelegate(chain.LatestLockNumber, chain.NodeConfigs[1].OperatorAddress, walletName)

	// create a text prop, deposit and vote yes
	node.SubmitTextProposal("superfluid vote overwrite test", sdk.NewCoin(appparams.BaseCoinUnit, sdk.NewInt(config.InitialMinDeposit)), false)
	chain.LatestProposalNumber += 1
	node.DepositProposal(chain.LatestProposalNumber, false)
	for _, node := range chain.NodeConfigs {
		node.VoteYesProposal(initialization.ValidatorWalletName, chain.LatestProposalNumber)
	}

	// set delegator vote to no
	node.VoteNoProposal(walletName, chain.LatestProposalNumber)

	s.Eventually(
		func() bool {
			noTotal, yesTotal, noWithVetoTotal, abstainTotal, err := node.QueryPropTally(chain.LatestProposalNumber)
			if err != nil {
				return false
			}
			if abstainTotal.Int64()+noTotal.Int64()+noWithVetoTotal.Int64()+yesTotal.Int64() <= 0 {
				return false
			}
			return true
		},
		1*time.Minute,
		10*time.Millisecond,
		"Osmosis node failed to retrieve prop tally",
	)
	noTotal, _, _, _, _ := node.QueryPropTally(chain.LatestProposalNumber)
	noTotalFinal, err := strconv.Atoi(noTotal.String())
	s.NoError(err)

	s.Eventually(
		func() bool {
			intAccountBalance, err := node.QueryIntermediaryAccount("gamm/pool/1", chain.NodeConfigs[1].OperatorAddress)
			s.Require().NoError(err)
			if err != nil {
				return false
			}
			if noTotalFinal != intAccountBalance {
				fmt.Printf("noTotalFinal %v does not match intAccountBalance %v", noTotalFinal, intAccountBalance)
				return false
			}
			return true
		},
		1*time.Minute,
		10*time.Millisecond,
		"superfluid delegation vote overwrite not working as expected",
	)
}

func (s *IntegrationTestSuite) TestStateSync() {
	if s.skipStateSync {
		s.T().Skip()
	}

	chain := s.configurer.GetChainConfig(0)
	runningNode, err := chain.GetDefaultNode()
	s.Require().NoError(err)

	persistenrPeers := chain.GetPersistentPeers()

	stateSyncHostPort := fmt.Sprintf("%s:26657", runningNode.Name)
	stateSyncRPCServers := []string{stateSyncHostPort, stateSyncHostPort}

	// get trust height and trust hash.
	trustHeight, err := runningNode.QueryCurrentHeight()
	s.Require().NoError(err)

	trustHash, err := runningNode.QueryHashFromBlock(trustHeight)
	s.Require().NoError(err)

	stateSynchingNodeConfig := &initialization.NodeConfig{
		Name:               "state-sync",
		Pruning:            "default",
		PruningKeepRecent:  "0",
		PruningInterval:    "0",
		SnapshotInterval:   1500,
		SnapshotKeepRecent: 2,
	}

	tempDir, err := os.MkdirTemp("", "osmosis-e2e-statesync-")
	s.Require().NoError(err)

	// configure genesis and config files for the state-synchin node.
	nodeInit, err := initialization.InitSingleNode(
		chain.Id,
		tempDir,
		filepath.Join(runningNode.ConfigDir, "config", "genesis.json"),
		stateSynchingNodeConfig,
		time.Duration(chain.VotingPeriod),
		//time.Duration(chain.ExpeditedVotingPeriod),
		trustHeight,
		trustHash,
		stateSyncRPCServers,
		persistenrPeers,
	)
	s.Require().NoError(err)

	stateSynchingNode := chain.CreateNode(nodeInit)

	// ensure that the running node has snapshots at a height > trustHeight.
	hasSnapshotsAvailable := func(syncInfo coretypes.SyncInfo) bool {
		snapshotHeight := runningNode.SnapshotInterval
		if uint64(syncInfo.LatestBlockHeight) < snapshotHeight {
			s.T().Logf("snapshot height is not reached yet, current (%d), need (%d)", syncInfo.LatestBlockHeight, snapshotHeight)
			return false
		}

		snapshots, err := runningNode.QueryListSnapshots()
		s.Require().NoError(err)

		for _, snapshot := range snapshots {
			if snapshot.Height > uint64(trustHeight) {
				s.T().Log("found state sync snapshot after trust height")
				return true
			}
		}
		s.T().Log("state sync snashot after trust height is not found")
		return false
	}
	runningNode.WaitUntil(hasSnapshotsAvailable)

	// start the state synchin node.
	err = stateSynchingNode.Run()
	s.NoError(err)

	// ensure that the state synching node cathes up to the running node.
	s.Require().Eventually(func() bool {
		stateSyncNodeHeight, err := stateSynchingNode.QueryCurrentHeight()
		s.NoError(err)

		runningNodeHeight, err := runningNode.QueryCurrentHeight()
		s.NoError(err)

		return stateSyncNodeHeight == runningNodeHeight
	},
		3*time.Minute,
		500*time.Millisecond,
	)

	// stop the state synching node.
	err = chain.RemoveNode(stateSynchingNode.Name)
	s.NoError(err)
}

func (s *IntegrationTestSuite) TestExpeditedProposals() {
	if !s.skipUpgrade {
		s.T().Skip("this can be re-enabled post v12")
	}

	chain := s.configurer.GetChainConfig(0)
	node, err := chain.GetDefaultNode()
	s.NoError(err)

	node.SubmitTextProposal("expedited text proposal", sdk.NewCoin(appparams.BaseCoinUnit, sdk.NewInt(config.InitialMinExpeditedDeposit)), true)
	chain.LatestProposalNumber += 1
	node.DepositProposal(chain.LatestProposalNumber, true)
	totalTimeChan := make(chan time.Duration, 1)
	go node.QueryPropStatusTimed(chain.LatestProposalNumber, "PROPOSAL_STATUS_PASSED", totalTimeChan)
	for _, node := range chain.NodeConfigs {
		node.VoteYesProposal(initialization.ValidatorWalletName, chain.LatestProposalNumber)
	}
	// if querying proposal takes longer than timeoutPeriod, stop the goroutine and error
	var elapsed time.Duration
	timeoutPeriod := time.Duration(2 * time.Minute)
	select {
	case elapsed = <-totalTimeChan:
	case <-time.After(timeoutPeriod):
		err := fmt.Errorf("go routine took longer than %s", timeoutPeriod)
		s.Require().NoError(err)
	}

	// compare the time it took to reach pass status to expected expedited voting period
	expeditedVotingPeriodDuration := time.Duration(chain.ExpeditedVotingPeriod * 1000000000)
	timeDelta := elapsed - expeditedVotingPeriodDuration
	// ensure delta is within one second of expected time
	s.Require().Less(timeDelta, time.Second)
	s.T().Logf("expeditedVotingPeriodDuration within one second of expected time: %v", timeDelta)
	close(totalTimeChan)
}
