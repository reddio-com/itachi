# itachi
a decentralized sequencer for Starknet

#### Build
```shell
git submodule init
git submodule update --recursive --checkout
make build
```

### Build genesis contract
```shell
python3 scripts/abi_dumps.py
```