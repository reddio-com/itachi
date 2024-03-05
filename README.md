# itachi
a decentralized sequencer for Starknet

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

### RPC
- [x] addDeclareTransaction
- [x] addDeployAccountTransaction
- [x] addInvokeTransaction
- [x] call
- [ ] estimateFee
- [x] getTransactionReceipt
- [x] getTransactionByHash
- [x] getNonce
- [x] getTransactionStatus
- [x] getClass
- [x] getClassAt
- [ ] getClassHashAt
- [ ] blockHashAndNumber
- [ ] chainId
- [ ] syncing
- [ ] getTransactionByBlockIdAndIndex
- [ ] getBlockTransactionCount
- [ ] estimateMessageFee
- [ ] blockNumber
- [ ] specVersion
- [ ] traceTransaction
- [ ] simulateTransactions
- [ ] traceBlockTransactions
- [x] getStorageAt
- [ ] getStateUpdate
