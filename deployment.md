# 部署链服务

## 下载代码
```
git clone git@github.com:mttdLabs/mtt-chain.git
cd mtt-chain
```

## 安装
### 运行安装脚本
```
./install.sh
```

### 检查是否安装成功
```
mttd
```
是否打印出如下内容
```
Start mttd node

Usage:
mttd [command]

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
-h, --help                     help for mttd
--home string              directory for config and data (default "/Users/zzz/.mttd")
--keyring-backend string   Select keyring's backend (default "os")
--log_format string        The logging format (json|plain) (default "plain")
--log_level string         The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
--node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
--trace                    print out full stack trace on errors

Use "mttd [command] --help" for more information about a command.
```

## 初始化
### 初始化管理员钱包
```
mttd keys add alice --keyring-backend file --home ~/.mtt
```
获取到模块管理员地址之后,更换init.sh里的app_state["erc20"]["params"]["admin"]字段
### 第一个节点
设置创世mint token的数量  
21e7+4e6+1e5+1e4   
总量：214110000
- 2.1亿用于设置挖矿奖励
- 400万 4个节点质押100万个
- 10万 用于旧链到新链的mtt代币迁移，上线之后这10w个会根据先前的快照数据进行代币分发
- 1万 用于开发（部署合约，调用桥的合约）
在上线完成之后，将相应数量的币打到桥合约地址即可，保证能够进行兑付，主网的桥合约可以沿用先前的
```
./init.sh
```
运行完成，确认无错误之后就可以开始运行节点了    
使用创世钱包往模块注资2.1亿代币
```
mttd tx erc20 fund 210000000 --from alice --home ~/.mtt/ --keyring-backend file --chain-id mtt_6118-1
```
挖矿暂时不会启动，设置的区块高度为1年后，等后面确定具体的日期后可以使用管理员钱包进行调整  
先使用创世钱包给管理员钱包转点gas费，然后使用管理员钱包设置桥地址，需要先部署桥合约获得桥地址
```
mttd tx erc20 set-bridge 0x2773D083ace5ad7a46E60477B02193E635F366E0 --from alice --home ~/.mtt --keyring-backend file
```

### 其他节点
```
./init.sh
```
需要从第一个节点的配置中拷贝genesis.json
```
cp genesis.json ~/.mtt/config
```
更改配置,其中node_id是第一个节点的node_id,可以通过
```
mttd tendermint show-node-id 
```
ip就是第一个节点的ip
```
vim ~/.mtt/config/config.toml
#persistent_peers = ""
persistent_peers = "[node_id]@[ip]:26656"
```

## 运行
```
mttd start --pruning=nothing --trace --json-rpc.api eth,txpool,net,web3,debug,trace --log_level info --home /home/ubuntu/.mttd --evm.tracer struct
```
