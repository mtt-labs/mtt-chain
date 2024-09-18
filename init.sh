KEY="alice"
CHAINID="mtt_6880-1"
MONIKER="mtt-node-one"
KEYRING="test"
LOGLEVEL="info"
HOMEDIR="/Users/zzz/.mtt"

# Reinstall daemon
rm -rf $HOMEDIR/data/
rm -rf $HOMEDIR/config/

# Set client config
mttd config keyring-backend $KEYRING --home $HOMEDIR
mttd config chain-id $CHAINID --home $HOMEDIR

# if $KEY exists it should be deleted
mttd keys add $KEY --keyring-backend $KEYRING --home $HOMEDIR

ADDR=$(mttd keys show alice | grep "address" | awk '{print $3}')

# Set moniker and chain-id for cosmos (Moniker can be anything, chain-id must be an integer)
mttd init $MONIKER --chain-id $CHAINID --home $HOMEDIR --overwrite

cp $HOMEDIR/config/genesis.json $HOMEDIR/config/tmp_genesis.json

# Set gas limit in genesis
cat $HOMEDIR/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="40000000"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json

# Change parameter amtt denominations to amtt
cat $HOMEDIR/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="amtt"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="amtt"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["gov"]["params"]["min_deposit"][0]["denom"]="mtt"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["gov"]["params"]["min_deposit"][0]["amount"]="10"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["inflation"]["params"]["mint_denom"]="amtt"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["mint"]["params"]["coin"]["denom"]="amtt"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["mint"]["params"]["blocks_per_year"]="10512000"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["mint"]["params"]["begin_block"]="6048000"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["erc20"]["params"]["admin"]="'$ADDR'"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
#cat $HOMEDIR/config/genesis.json | jq '.app_state["feemarket"]["params"]["min_gas_price"]="10000000000"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["distribution"]["params"]["community_tax"]="0.000000000000000000"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["mint"]["params"]["coin"]["amount"]="210000000000000000000000000"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json

#
# disable produce empty block
if [[ "$OSTYPE" == "darwin"* ]]; then
#    sed -i '' 's/127.0.0.1:26657/0.0.0.0:26657/g' $HOMEDIR/config/config.toml
    sed -i '' '$H;x;1,/enable = false/s/enable = false/enable = true/;1d' $HOMEDIR/config/app.toml
    sed -i '' ' s/swagger = false/swagger = true/g' $HOMEDIR/config/app.toml
    sed -i '' 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' $HOMEDIR/config/app.toml
    sed -i '' 's/minimum-gas-prices = ""/minimum-gas-prices = "0amtt"/g' $HOMEDIR/config/app.toml
    sed -i '' 's/localhost/0.0.0.0/g' $HOMEDIR/config/app.toml
    sed -i '' 's/127.0.0.1/0.0.0.0/g' $HOMEDIR/config/app.toml
    sed -i '' 's/localhost/0.0.0.0/g' $HOMEDIR/config/config.toml
    sed -i '' 's/127.0.0.1/0.0.0.0/g' $HOMEDIR/config/config.toml
  else
#    sed -i 's/127.0.0.1:26657/0.0.0.0:26657/g' $HOMEDIR/config/config.toml
    sed -i '$H;x;1,/enable = false/s/enable = false/enable = true/;1d' $HOMEDIR/config/app.toml
    sed -i ' s/swagger = false/swagger = true/g' $HOMEDIR/config/app.toml
    sed -i 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' $HOMEDIR/config/app.toml
    sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "0amtt"/g' $HOMEDIR/config/app.toml
    sed -i 's/localhost/0.0.0.0/g' $HOMEDIR/config/app.toml
    sed -i 's/127.0.0.1/0.0.0.0/g' $HOMEDIR/config/app.toml
    sed -i 's/localhost/0.0.0.0/g' $HOMEDIR/config/config.toml
    sed -i 's/127.0.0.1/0.0.0.0/g' $HOMEDIR/config/config.toml
fi

if [[ "$OSTYPE" == "darwin"* ]]; then
      sed -i '' 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOMEDIR/config/config.toml
      sed -i '' 's/timeout_propose = "3s"/timeout_propose = "1800ms"/g' $HOMEDIR/config/config.toml
      sed -i '' 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "500ms"/g' $HOMEDIR/config/config.toml
      sed -i '' 's/timeout_prevote = "1s"/timeout_prevote = "600ms"/g' $HOMEDIR/config/config.toml
      sed -i '' 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "500ms"/g' $HOMEDIR/config/config.toml
      sed -i '' 's/timeout_precommit = "1s"/timeout_precommit = "600ms"/g' $HOMEDIR/config/config.toml
      sed -i '' 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "500ms"/g' $HOMEDIR/config/config.toml
      sed -i '' 's/timeout_commit = "5s"/timeout_commit = "3s"/g' $HOMEDIR/config/config.toml
      sed -i '' 's/timeout_broadcast_tx_commit = "2m30s"/timeout_broadcast_tx_commit = "10s"/g' $HOMEDIR/config/config.toml
  else
      sed -i 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOMEDIR/config/config.toml
      sed -i 's/timeout_propose = "3s"/timeout_propose = "1800ms"/g' $HOMEDIR/config/config.toml
      sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "500ms"/g' $HOMEDIR/config/config.toml
      sed -i 's/timeout_prevote = "1s"/timeout_prevote = "600ms"/g' $HOMEDIR/config/config.toml
      sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "500ms"/g' $HOMEDIR/config/config.toml
      sed -i 's/timeout_precommit = "1s"/timeout_precommit = "600ms"/g' $HOMEDIR/config/config.toml
      sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "500ms"/g' $HOMEDIR/config/config.toml
      sed -i 's/timeout_commit = "5s"/timeout_commit = "3s"/g' $HOMEDIR/config/config.toml
      sed -i 's/timeout_broadcast_tx_commit = "2m30s"/timeout_broadcast_tx_commit = "10s"/g' $HOMEDIR/config/config.toml
fi

# Allocate genesis accounts (cosmos formatted addresses)
mttd add-genesis-account $KEY 214110000000000000000000000amtt --keyring-backend $KEYRING --home $HOMEDIR

# Update total supply with claim values
validators_supply=$(cat $HOMEDIR/config/genesis.json | jq -r '.app_state["bank"]["supply"][0]["amount"]')
total_supply=214110000000000000000000000
cat $HOMEDIR/config/genesis.json | jq -r '.app_state["bank"]["supply"][0]["denom"]="amtt"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq -r --arg total_supply "$total_supply" '.app_state["bank"]["supply"][0]["amount"]=$total_supply' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json

# Sign genesis transaction
mttd gentx $KEY 1000000000000000000000000amtt --keyring-backend $KEYRING --chain-id $CHAINID --home $HOMEDIR --fees 20000000000000000amtt

# Collect genesis tx
mttd collect-gentxs --home $HOMEDIR

# Run this to ensure everything worked and that the genesis file is setup correctly
mttd validate-genesis --home $HOMEDIR

#Start the node
echo mttd start --pruning=nothing --json-rpc.api eth,txpool,net,web3,trace,debug --log_level info --home $HOMEDIR
