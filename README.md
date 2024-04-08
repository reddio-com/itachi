# itachi
A decentralized sequencer for Starknet.   

## Overall Structure
![image](docs/images/itachi_arch.png)

## Build & Run
### Prerequisites
- rustc 1.74.0 (79e9716c9 2023-11-13)  
- go 1.21

### Docker Pull & Run
```shell
docker pull ghcr.io/reddio-com/itachi:{version}
docker-compose up
```

### Source code Build & Run
```shell
git submodule init
git submodule update --recursive --checkout
make build

./itachi
```

### Reset Chain
Reset Chain will clean all the stored history data locally. 
```shell
make reset
```  
## Test
### Use [Starkli](https://github.com/xJonathanLEI/starkli)  


### Use [starknet.py](https://github.com/software-mansion/starknet.py)
You can use our python demo: https://github.com/reddio-com/itachi-testing


## Configs  
### Chain Configs
The default config file of Itachi chain is `conf/cairo_cfg.toml`
### Genesis Configs
The genesis configs of Itachi chain is same as Madara. You can learn more details by [docs](docs/genesis.md)


## Starknet RPC
#### Compatible Versions: 
- 0.5.1
- 0.6.0   
#### Compatible RPC methods
- [x] addDeclareTransaction
- [x] addDeployAccountTransaction
- [x] addInvokeTransaction
- [x] call
- [x] estimateFee
- [x] getTransactionReceipt
- [x] getTransactionByHash
- [x] getNonce
- [x] getTransactionStatus
- [x] getClass
- [x] getClassAt
- [x] getClassHashAt
- [ ] blockHashAndNumber
- [x] chainId
- [ ] syncing
- [ ] getTransactionByBlockIdAndIndex
- [ ] getBlockTransactionCount
- [ ] estimateMessageFee
- [ ] blockNumber
- [x] specVersion
- [ ] traceTransaction
- [x] simulateTransactions
- [ ] traceBlockTransactions
- [x] getStorageAt
- [ ] getStateUpdate
