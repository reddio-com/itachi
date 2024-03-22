# itachi
A decentralized sequencer for Starknet

### Prerequisites
- rustc 1.74.0 (79e9716c9 2023-11-13)  
- go 1.21.4

### Build Code
```shell
git submodule init
git submodule update --recursive --checkout
make build
```

### Reset Chain
```shell
make reset
```

### Build genesis contract
```shell
python3 scripts/abi_dumps.py
```

### Starknet RPC
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
- [ ] specVersion
- [ ] traceTransaction
- [x] simulateTransactions
- [ ] traceBlockTransactions
- [x] getStorageAt
- [ ] getStateUpdate
