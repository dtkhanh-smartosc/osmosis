package keeper_test

import (
	"encoding/hex"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/suite"

	"github.com/osmosis-labs/osmosis/v23/app/apptesting"
	"github.com/osmosis-labs/osmosis/v23/x/authenticator/authenticator"
	"github.com/osmosis-labs/osmosis/v23/x/authenticator/testutils"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper
	am *authenticator.AuthenticatorManager
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	s.Reset()
	s.am = authenticator.NewAuthenticatorManager()

	// Register the SigVerificationAuthenticator
	s.am.InitializeAuthenticators([]authenticator.Authenticator{
		authenticator.SignatureVerificationAuthenticator{},
		testutils.TestingAuthenticator{
			Approve:        testutils.Always,
			GasConsumption: 10,
			Confirm:        testutils.Always,
		},
	})
}

func (s *KeeperTestSuite) TestKeeper_AddAuthenticator() {
	ctx := s.Ctx

	// Set up account
	key := "6cf5103c60c939a5f38e383b52239c5296c968579eec1c68a47d70fbf1d19159"
	bz, _ := hex.DecodeString(key)
	priv := &secp256k1.PrivKey{Key: bz}
	accAddress := sdk.AccAddress(priv.PubKey().Address())

	err := s.App.AuthenticatorKeeper.AddAuthenticator(
		ctx,
		accAddress,
		"SignatureVerificationAuthenticator",
		priv.PubKey().Bytes(),
	)
	s.Require().NoError(err, "Should successfully add a SignatureVerificationAuthenticator")

	err = s.App.AuthenticatorKeeper.AddAuthenticator(
		ctx,
		accAddress,
		"MessageFilterAuthenticator",
		[]byte(`{"@type":"/cosmos.bank.v1beta1.MsgSend"}`),
	)
	s.Require().NoError(err, "Should successfully add a MessageFilterAuthenticator")

	err = s.App.AuthenticatorKeeper.AddAuthenticator(
		ctx,
		accAddress,
		"SignatureVerificationAuthenticator",
		[]byte("BrokenBytes"),
	)
	s.Require().Error(err, "Should have failed as OnAuthenticatorAdded fails")

	s.App.AuthenticatorManager.ResetAuthenticators()
	err = s.App.AuthenticatorKeeper.AddAuthenticator(
		ctx,
		accAddress,
		"MessageFilterAuthenticator",
		[]byte(`{"@type":"/cosmos.bank.v1beta1.MsgSend"`),
	)
	s.Require().Error(err, "Authenticator not registered so should fail")
}

func (s *KeeperTestSuite) TestKeeper_GetAuthenticatorDataForAccount() {
	ctx := s.Ctx

	// Set up account
	key := "6cf5103c60c939a5f38e383b52239c5296c968579eec1c68a47d70fbf1d19159"
	bz, _ := hex.DecodeString(key)
	priv := &secp256k1.PrivKey{Key: bz}
	accAddress := sdk.AccAddress(priv.PubKey().Address())

	err := s.App.AuthenticatorKeeper.AddAuthenticator(
		ctx,
		accAddress,
		"SignatureVerificationAuthenticator",
		priv.PubKey().Bytes(),
	)
	s.Require().NoError(err, "Should successfully add a SignatureVerificationAuthenticator")

	err = s.App.AuthenticatorKeeper.AddAuthenticator(
		ctx,
		accAddress,
		"SignatureVerificationAuthenticator",
		priv.PubKey().Bytes(),
	)
	s.Require().NoError(err, "Should successfully add a MessageFilterAuthenticator")

	authenticators, err := s.App.AuthenticatorKeeper.GetAuthenticatorDataForAccount(ctx, accAddress)
	s.Require().NoError(err)
	s.Require().Equal(len(authenticators), 2, "Getting authenticators returning incorrect data")
}

func (s *KeeperTestSuite) TestKeeper_GetAndSetAuthenticatorId() {
	ctx := s.Ctx

	authenticatorId := s.App.AuthenticatorKeeper.GetNextAuthenticatorId(ctx)
	s.Require().Equal(authenticatorId, uint64(0), "Get authenticator id returned incorrect id")

	authenticatorId = s.App.AuthenticatorKeeper.GetNextAuthenticatorId(ctx)
	s.Require().Equal(authenticatorId, uint64(0), "Get authenticator id returned incorrect id")

	// Set up account
	key := "6cf5103c60c939a5f38e383b52239c5296c968579eec1c68a47d70fbf1d19159"
	bz, _ := hex.DecodeString(key)
	priv := &secp256k1.PrivKey{Key: bz}
	accAddress := sdk.AccAddress(priv.PubKey().Address())

	err := s.App.AuthenticatorKeeper.AddAuthenticator(
		ctx,
		accAddress,
		"SignatureVerificationAuthenticator",
		priv.PubKey().Bytes(),
	)
	s.Require().NoError(err, "Should successfully add a SignatureVerificationAuthenticator")

	authenticatorId = s.App.AuthenticatorKeeper.GetNextAuthenticatorId(ctx)
	s.Require().Equal(authenticatorId, uint64(1), "Get authenticator id returned incorrect id")

	err = s.App.AuthenticatorKeeper.AddAuthenticator(
		ctx,
		accAddress,
		"SignatureVerificationAuthenticator",
		priv.PubKey().Bytes(),
	)
	s.Require().NoError(err, "Should successfully add a MessageFilterAuthenticator")

	authenticatorId = s.App.AuthenticatorKeeper.GetNextAuthenticatorId(ctx)
	s.Require().Equal(authenticatorId, uint64(2), "Get authenticator id returned incorrect id")
}

func (s *KeeperTestSuite) TestKeeper_GetSelectedAuthenticatorForAccount() {
	ctx := s.Ctx

	// Set up account
	key := "6cf5103c60c939a5f38e383b52239c5296c968579eec1c68a47d70fbf1d19159"
	bz, _ := hex.DecodeString(key)
	priv := &secp256k1.PrivKey{Key: bz}
	accAddress := sdk.AccAddress(priv.PubKey().Address())

	err := s.App.AuthenticatorKeeper.AddAuthenticator(
		ctx,
		accAddress,
		"SignatureVerificationAuthenticator",
		priv.PubKey().Bytes(),
	)
	s.Require().NoError(err, "Should successfully add a SignatureVerificationAuthenticator")

	err = s.App.AuthenticatorKeeper.AddAuthenticator(
		ctx,
		accAddress,
		"MessageFilterAuthenticator",
		[]byte(`{"@type":"/cosmos.bank.v1beta1.MsgSend"}`),
	)
	s.Require().NoError(err, "Should successfully add a MessageFilterAuthenticator")

	// Test getting a selected authenticator from the store
	selectedAuthenticator, err := s.App.AuthenticatorKeeper.GetInitializedAuthenticatorForAccount(ctx, accAddress, 1)
	s.Require().NoError(err)
	s.Require().Equal(selectedAuthenticator.Authenticator.Type(), "MessageFilterAuthenticator", "Getting authenticators returning incorrect data")

	selectedAuthenticator, err = s.App.AuthenticatorKeeper.GetInitializedAuthenticatorForAccount(ctx, accAddress, 2)
	s.Require().NoError(err)
	s.Require().Equal(selectedAuthenticator.Authenticator.Type(), "SignatureVerificationAuthenticator", "Getting authenticators returning incorrect data")
	s.Require().Equal(selectedAuthenticator.Id, uint64(0), "Incorrect ID returned from store")

	err = s.App.AuthenticatorKeeper.AddAuthenticator(
		ctx,
		accAddress,
		"MessageFilterAuthenticator",
		[]byte(`{"@type":"/cosmos.bank.v1beta1.MsgSend"}`),
	)
	s.Require().NoError(err, "Should successfully add a MessageFilterAuthenticator")

	// Remove a registered authenticator from the authenticator manager
	ar := s.App.AuthenticatorManager.GetRegisteredAuthenticators()
	for _, a := range ar {
		if a.Type() == "MessageFilterAuthenticator" {
			s.App.AuthenticatorManager.UnregisterAuthenticator(authenticator.MessageFilterAuthenticator{})
		}
	}

	// Try to get an authenticator that has been removed from the store
	selectedAuthenticator, err = s.App.AuthenticatorKeeper.GetInitializedAuthenticatorForAccount(ctx, accAddress, 2)
	s.Require().Error(err)
	s.Require().ErrorContains(err, "authenticator not registered in manager")

	// Reset the authenticator manager to see how GetInitializedAuthenticatorForAccount behaves
	s.App.AuthenticatorManager.ResetAuthenticators()
	selectedAuthenticator, err = s.App.AuthenticatorKeeper.GetInitializedAuthenticatorForAccount(ctx, accAddress, 3)
	s.Require().NoError(err)
	s.Require().Equal(selectedAuthenticator.Id, uint64(0), "Incorrect ID returned from store")
	s.Require().Equal(selectedAuthenticator.Authenticator, nil, "Returned authenticator from store but nothing registered in manager")
}
