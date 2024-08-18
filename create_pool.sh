TX_FLAGS="--node http://localhost:26657 --keyring-backend test --from deployer --chain-id localosmosis --gas-prices 0.1uosmo --gas auto --gas-adjustment 1.3 --yes"


deployer=$(osmosisd keys show deployer --address --keyring-backend test)

osmosisd tx gamm create-pool --pool-file create_pool.json $TX_FLAGS