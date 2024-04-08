# itachi
A decentralized sequencer for Starknet.   

![image](docs/images/itachi_arch.png)

### Prerequisites
- rustc 1.74.0 (79e9716c9 2023-11-13)  
- go 1.21.4

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

### Build genesis contract
It will refactor the cairo1 json files for itachi to load.
If you want to add new cairo1 contracts for genesis, you should run this script.
```shell
python3 scripts/abi_dumps.py
```

### Genesis Configs
The genesis configs of Itachi is same as Madara. You can learn more details by [docs](docs/genesis.md)

### Starknet RPC
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
