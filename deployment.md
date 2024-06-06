# 部署链服务

## 下载代码
```
git clone git@github.com:mttdLabs/mttd-chain.git
cd mttd-chain
```

## 安装
### 运行安装脚本
```
./install.sh
```

### 检查是否安装成功
```
mttdd
```
是否打印出如下内容
```
Start mttd node

Usage:
mttdd [command]

Available Commands:
add-genesis-account Add a genesis account to genesis.json
collect-gentxs      Collect genesis txs and output a genesis.json file
config              Create or query an application CLI configuration file
debug               Tool for helping with debugging your application
export              Export state to JSON
gentx               Generate a genesis tx carrying a self delegation
help                Help about any command
index-eth-tx        Index historical eth txs
init                Initialize private validator, p2p, genesis, and application configuration files
keys                Manage your application's keys
migrate             Migrate genesis to a specified target version
query               Querying subcommands
rollback            rollback cosmos-sdk and tendermint state by one height
rosetta             spin up a rosetta server
start               Run the full node
status              Query remote node for status
tendermint          Tendermint subcommands
testnet             subcommands for starting or configuring local testnets
tx                  Transactions subcommands
validate-genesis    validates the genesis file at the default location or at the location passed as an arg
version             Print the application binary version information

Flags:
-b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
--chain-id string          Specify Chain ID for sending Tx (default "testnet")
--fees string              Fees to pay along with transaction; eg: 10aphoton
--from string              Name or address of private key with which to sign
--gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
--gas-prices string        Gas prices to determine the transaction fee (e.g. 10aphoton)
-h, --help                     help for mttdd
--home string              directory for config and data (default "/Users/zzz/.mttd")
--keyring-backend string   Select keyring's backend (default "os")
--log_format string        The logging format (json|plain) (default "plain")
--log_level string         The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
--node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
--trace                    print out full stack trace on errors

Use "mttdd [command] --help" for more information about a command.
```

## 初始化
### 第一个节点
```
./init.sh
```
运行完成，确认无错误之后就可以开始运行节点了

### 其他节点
```
./init.sh
```
需要从第一个节点的配置中拷贝genesis.json
```
cp genesis.json ~/.mttd/config
```
更改配置,其中node_id是第一个节点的node_id,可以通过
```
mttdd tendermint show-node-id 
```
ip就是第一个节点的ip
```
vim ~/.mttd/config/config.toml
#persistent_peers = ""
persistent_peers = "[node_id]@[ip]:26656"
```

## 运行
```
mttdd start --pruning=nothing --trace --json-rpc.api eth,txpool,net,web3,debug,trace --log_level info --home /home/ubuntu/.mttd --evm.tracer struct
```
