# Genesis

The genesis of the chain can be found in the [configs]
(<https://github.com/reddio-com/itachi/tree/main/conf/genesis>)
folder. The genesis is defined in the form of a JSON file containing the
following:

- contract_classes: list of maps containing the class hash and the class.
- contracts: list of maps containing the contract address and the associated
  class hash.
- storage: list of tuples containing the storage key and the storage value.
  Please note that the storage key is itself a tuple, containing the contract
  address for which storage is set and the
  [Starknet storage key](https://docs.starknet.io/documentation/architecture_and_concepts/Smart_Contracts/contract-storage/#storage_variables).

The below defines all hardcoded values set in the geneses:

## Node genesis [link](https://github.com/keep-starknet-strange/madara/tree/main/configs/genesis-assets/genesis.json)

### Contract classes node genesis

<!-- markdownlint-disable MD013 -->

| Class hash                                                         | Definition                                     |
| :----------------------------------------------------------------- | :--------------------------------------------- |
| 0x0279d77db761fba82e0054125a6fdb5f6baa6286fa3fb73450cc44d193c2d37f | No validation account class hash               |
| 0x02f99bf9799ada84cd5ac0d0fe36b9d8f65efcb377cd2e8cf8309ad2daf15e4b | No validation account class hash cairo 1       |
| 0x06f0d6f6ae72e1a507ff4b65181291642889742dbf8f1a53e9ec1c595d01ba7d | Argent account class hash                      |
| 0x0424b7f61e3c5dfd74400d96fdea7e1f0bf2757f31df04387eaa957f095dd7b9 | Proxy class hash                               |
| 0x2c2b8f559e1221468140ad7b2352b1a5be32660d0bf1a3ae3a054a4ec5254e4  | Braavos account class hash                     |
| 0x5aa23d5bb71ddaa783da7ea79d405315bafa7cf0387a74f4593578c3e9e6570  | Braavos account base implementation class hash |
| 0x07db5c2c2676c2a5bfc892ee4f596b49514e3056a0eee8ad125870b4fb1dd909 | Braavos account call aggregator class hash     |
| 0x3131fa018d520a037686ce3efddeab8f28895662f019ca3ca18a626650f7d1e  | Proxy class hash                               |
| 0x006280083f8c2a2db9f737320d5e3029b380e0e820fe24b8d312a6a34fdba0cd | Openzeppelin account class hash                |
| 0x04c5efa8dc6f0554da51f125d04e379ac41153a8b837391083a8dc3771a33388 | Test contract class hash                       |
| 0x0372ee6669dc86563007245ed7343d5180b96221ce28f44408cff2898038dbd4 | ERC20 class hash                               |
| 0x077cc28ed3c661419fda16bf120fb81f1f8f28617f5543b05a86d63b0926bbf4 | ERC721 class hash                              |
| 0x07b3e05f48f0c69e4a65ce5e076a66271a527aff2c34ce1083ec6e1526997a69 | Universal deployer class hash                  |
| 0x01a736d6ed154502257f02b1ccdf4d9d1089f80811cd6acad48e6b6a9d1f2003 | Argent Cairo 1 Account                         |
| 0x04c6d6cf894f8bc96bb9c525e6853e5483177841f7388f74a46cfda6f028c755 | OpenZepplin Cairo 1 Account                    |

<!-- markdownlint-disable MD013 -->

### Predeployed accounts node genesis

<!-- markdownlint-disable MD013 -->

| Contract address | Class hash                                                         | Name                            | Optional private key |
| :--------------- | :----------------------------------------------------------------- | :------------------------------ | :------------------- |
| 0x1              | 0x0279d77db761fba82e0054125a6fdb5f6baa6286fa3fb73450cc44d193c2d37f | No Validation Account           | null                 |
| 0x2              | 0x06f0d6f6ae72e1a507ff4b65181291642889742dbf8f1a53e9ec1c595d01ba7d | Argent Account                  | `pk`                 |
| 0x3              | 0x006280083f8c2a2db9f737320d5e3029b380e0e820fe24b8d312a6a34fdba0cd | Openzeppelin Account            | `pk`                 |
| 0x4              | 0x02f99bf9799ada84cd5ac0d0fe36b9d8f65efcb377cd2e8cf8309ad2daf15e4b | No Validation Account (cairo 1) | null                 |

where `pk` is the following vector of `u8`:

```rust
[48,120,48,48,99,49,99,102,49,52,57,48,100,101,49,51,53,50,56,54,53,51,48,49,98,98,56,55,48,53,49,52,51,102,51,101,102,57,51,56,102,57,55,102,100,102,56,57,50,102,49,48,57,48,100,99,98,53,97,99,55,98,99,100,49,100]
```

<!-- markdownlint-disable MD013 -->

### Contracts node genesis

<!-- markdownlint-disable MD013 -->

| Contract address                                                   | Class hash                                                         |
| :----------------------------------------------------------------- | :----------------------------------------------------------------- |
| 0x1                                                                | 0x0279d77db761fba82e0054125a6fdb5f6baa6286fa3fb73450cc44d193c2d37f |
| 0x2                                                                | 0x06f0d6f6ae72e1a507ff4b65181291642889742dbf8f1a53e9ec1c595d01ba7d |
| 0x3                                                                | 0x006280083f8c2a2db9f737320d5e3029b380e0e820fe24b8d312a6a34fdba0cd |
| 0x4                                                                | 0x02f99bf9799ada84cd5ac0d0fe36b9d8f65efcb377cd2e8cf8309ad2daf15e4b |
| 0x1111                                                             | 0x04c5efa8dc6f0554da51f125d04e379ac41153a8b837391083a8dc3771a33388 |
| 0x040e59c2c182a58fb0a74349bfa4769cbbcba32547591dd3fb1def8623997d00 | 0x0372ee6669dc86563007245ed7343d5180b96221ce28f44408cff2898038dbd4 |
| 0x040e59c2c182a58fb0a74349bfa4769cbbcba32547591dd3fb1def8623997d02 | 0x077cc28ed3c661419fda16bf120fb81f1f8f28617f5543b05a86d63b0926bbf4 |
| 0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7  | 0x0372ee6669dc86563007245ed7343d5180b96221ce28f44408cff2898038dbd4 |
| 0x041a78e741e5af2fec34b695679bc6891742439f7afb8484ecd7766661ad02bf | 0x07b3e05f48f0c69e4a65ce5e076a66271a527aff2c34ce1083ec6e1526997a69 |

<!-- markdownlint-disable MD013 -->

### Storage node genesis

The node storage is prefilled using the genesis in order to allow access to
prefunded accounts. Available accounts with unlimited funds are 0x1, 0x2, 0x3
and 0x4 (hence why the storage value we write is
0xffffffffffffffffffffffffffffffff for U256 low and U256 high).

Additionally, a public key
(0x3603a2692a2ae60abb343e832ee53b55d6b25f02a3ef1565ec691edc7a209b2) is set for
accounts 0x2 and 0x3, for which the signature will be checked during the
validation phases of the execution. Accounts 0x1 and 0x4 include an empty
validation phases, meaning no signature is required to execute any transactions
going through them.

Finally, 0x1 is set as the contract owner of contract
0x040e59c2c182a58fb0a74349bfa4769cbbcba32547591dd3fb1def8623997d02, which is
deployed as a ERC721 contract (given the class hash of 0x80000).

<!-- markdownlint-disable MD013 -->

| Contract address                                                   | Storage key                                                                                  | Storage value                                                                         |
| :----------------------------------------------------------------- | :------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| 0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7  | 0x341c1bdfd89f69748aa00b5742b03adbffd79b8e80cab5c50d91cd8c2a79be1 (ERC20_name)               | 0x4574686572                                                                          |
| 0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7  | 0x0b6ce5410fca59d078ee9b2a4371a9d684c530d697c64fbef0ae6d5e8f0ac72 (ERC20_symbol)             | 0x455448                                                                              |
| 0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7  | 0x1f0d4aa99431d246bac9b8e48c33e888245b15e9678f64f9bdfc8823dc8f979 (ERC20_decimals)           | 0x12                                                                                  |
| 0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7  | 0x7b62949c85c6af8a50c11c22927f9302f7a2e40bc93b4c988415915b0f97f09 (ERC20_balances(0x1).low)  | 0xffffffffffffffffffffffffffffffff (U128::MAX)                                        |
| 0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7  | 0x7b62949c85c6af8a50c11c22927f9302f7a2e40bc93b4c988415915b0f97f0a (ERC20_balances(0x1).high) | 0xffffffffffffffffffffffffffffffff (U128::MAX)                                        |
| 0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7  | 0x1d8bbc4f93f5ab9858f6c0c0de2769599fb97511503d5bf2872ef6846f2146f (ERC20_balances(0x2).low)  | 0xffffffffffffffffffffffffffffffff (U128::MAX)                                        |
| 0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7  | 0x1d8bbc4f93f5ab9858f6c0c0de2769599fb97511503d5bf2872ef6846f21470 (ERC20_balances(0x2).high) | 0xffffffffffffffffffffffffffffffff (U128::MAX)                                        |
| 0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7  | 0x262e096a838c0d8f34f641ff917d47d7dcb345c69efe61d9ab6b675e7340fc6 (ERC20_balances(0x3).low)  | 0xffffffffffffffffffffffffffffffff (U128::MAX)                                        |
| 0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7  | 0x262e096a838c0d8f34f641ff917d47d7dcb345c69efe61d9ab6b675e7340fc7 (ERC20_balances(0x3).high) | 0xffffffffffffffffffffffffffffffff (U128::MAX)                                        |
| 0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7  | 0x45abe05a3e7fb0c2ae1fa912be22a7dbc4832915e00562e2783dee710b9e4bc (ERC20_balances(0x4).low)  | 0xffffffffffffffffffffffffffffffff (U128::MAX)                                        |
| 0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7  | 0x45abe05a3e7fb0c2ae1fa912be22a7dbc4832915e00562e2783dee710b9e4bd (ERC20_balances(0x4).high) | 0xffffffffffffffffffffffffffffffff (U128::MAX)                                        |
| 0x040e59c2c182a58fb0a74349bfa4769cbbcba32547591dd3fb1def8623997d00 | 0x7b62949c85c6af8a50c11c22927f9302f7a2e40bc93b4c988415915b0f97f09 (ERC20_balances(0x1).low)  | 0xffffffffffffffffffffffffffffffff (U128::MAX)                                        |
| 0x040e59c2c182a58fb0a74349bfa4769cbbcba32547591dd3fb1def8623997d00 | 0x7b62949c85c6af8a50c11c22927f9302f7a2e40bc93b4c988415915b0f97f0a (ERC20_balances(0x1).high) | 0xffffffffffffffffffffffffffffffff (U128::MAX)                                        |
| 0x2                                                                | 0x1ccc09c8a19948e048de7add6929589945e25f22059c7345aaf7837188d8d05 (\_signer)                 | 0x3603a2692a2ae60abb343e832ee53b55d6b25f02a3ef1565ec691edc7a209b2 (Signer public key) |
| 0x3                                                                | 0x1379ac0624b939ceb9dede92211d7db5ee174fe28be72245b0a1a2abd81c98f (Account_public_key)       | 0x3603a2692a2ae60abb343e832ee53b55d6b25f02a3ef1565ec691edc7a209b2 (Signer public key) |
| 0x040e59c2c182a58fb0a74349bfa4769cbbcba32547591dd3fb1def8623997d02 | 0x2bd557f4ba80dfabefabe45e9b2dd35db1b9a78e96c72bc2b69b655ce47a930 (Ownable_owner)            | 0x1 (Owner)                                                                           |

<!-- markdownlint-disable MD013 -->