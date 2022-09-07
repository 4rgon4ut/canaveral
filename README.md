![image](https://user-images.githubusercontent.com/59182467/188896738-ed1c1eae-08c4-4579-afe7-a6ebc35c1be3.png)


# Welcome to Canaveral!
Canaveral is a cli tool written in Go for compiling, deploying and making calls to solidity smart contracts on Evmos blockchain
(potentially on every EVM-compatible chain).

Another words canaveral is a launch site for smart contracts.

- [Pre-Requisites](#pre-requisites)
- [Dependencies and Makefile](#dependencies-and-makefile)
- [Installation](#installation)
- [CLI](#cli)
- [Methods](#methods)
- [Further Scope](#further-scope)

## Pre-Requisites
- [Go v1.18+](https://go.dev/)
- [Solc](https://docs.soliditylang.org/en/v0.8.15/installing-solidity.html#macos-packages)
- [Abigen](https://goethereumbook.org/smart-contract-compile/)
- [Golangci-lint](https://golangci-lint.run/usage/install/#local-installation)
- [Ginkgo](https://onsi.github.io/ginkgo/#installing-ginkgo)

## Dependencies and Makefile
  `solc`, `abigen`, `golangci-lint` and `ginkgo` can be installed along with other dependencies by:
  ```
  make deps
  ```
  To get full list of available make commands:
  ```
  make help
  ```

## Installation
### 1. Firstly we need to configure Evmos local node.
There are several ways to do that. You can find more detailed information about every way in [docs](https://docs.evmos.org/developers/localnet/single_node.html).

I personally prefer using docker container:
```sh
docker run -it  -p 26657:26657 -p 26656:26656  -p 8545:8545  tharsishq/evmos:v8.0.0 bash
```
Configuration either be done [manually](https://docs.evmos.org/validators/quickstart/run_node.html#manual-deployment)
or using the `start-node.sh` script from
[canaveral](https://github.com/4rgon4ut/canaveral/blob/develop/start-node.sh) repo.

To use `start-node.sh` from evmos docker container:
```
docker cp path/to/start-node.sh <container_id>:/target/dir
```
And then just run:
```
sh ./start-node.sh
```

Or you can run commands from `start-node.sh` one by one manually:
```sh
# env vars
KEY="mykey"
CHAINID="evmos_9000-4"
MONIKER="localtestnet"
KEYRING="os" # for me only os works properly ('test' is not)
LOGLEVEL="info"

# clean evmos directory
rm -rf ~/.evmosd/

# init genesis files
evmosd init test --chain-id=$CHAINID

# generate keypair
evmosd keys add $KEY --keyring-backend=$KEYRING

# predefine validator balance and stake
evmosd add-genesis-account $KEY 1000000000000000000000aphoton,10000000000000000000aevmos,1000000000000000000000stake --chain-id=$CHAINID
evmosd gentx $KEY 1000000000000000000000stake --chain-id=$CHAINID --moniker=$MONIKER --keyring-backend=$KEYRING

evmosd collect-gentxs
evmosd validate-genesis

# start node
evmosd start  --log_level $LOGLEVEL  --json-rpc.api eth,txpool,personal,net,debug,web3
```
### 2. Add PrivateKey to Canaveral Config
To get available key pairs run:
```sh
evmosd keys list
```


Then print private key to terminal stdout:
```sh
evmosd keys unsafe-export-eth-key $KEYNAME --keyring-backend test


693F03A42E6F377D2305CB036EAE9BACCC09B230041CC786252A3BD5C34ED0FA
```
> :warning: **THIS IS UNSAFE AND ONLY CAN BE USED FOR LOCAL ENVIRONMENT!!!**

Add PrivateKey as `./config/config.toml` corresponding filed value
```toml
abi_dir = 'artifacts/abi'
bin_dir = 'artifacts/bin'
binds_dir = 'bindings'
contracts_dir = 'contracts'

--------------------------------------------------------------------------------
private_key = '693F03A42E6F377D2305CB036EAE9BACCC09B230041CC786252A3BD5C34ED0FA'
--------------------------------------------------------------------------------

rpc_addr = 'http://0.0.0.0'
rpc_port = '8545'

log_level = ''
```

### 3. Install CLI binary to your local go/bin
```
make install
```


## CLI
As I mentioned earlier Canaveral is a cli tool, so all interaction goes by specific commands.

To see full list of commands run:
```sh
canaveral help
```

### Methods

Canaveral allows user to compile, deploy and call every smart contract which is valid and can be compiled by `solc`.

```sh
canaveral compile [contract-name]

canaveral deploy [contract-name] arg1, arg2, ...

canaveral call [contract_name] [method] arg1, arg2, ...
```

### Constraints
Deploying contract constructor arguments and function calls arguments should be simple. This is because of casting CMD inputs(of type string) to specific solidity types.
I'm planning to spread supported types during time.

List of supported constructor types:
```
|    Go          |    Solidity    |
|                |                |
|----------------|----------------|
|    uint8       |    uint8       |
|----------------|----------------|
|    uint16      |    uint16      |
|----------------|----------------|
|    uint32      |    uint32      |
|----------------|----------------|
|    uint64      |    uint64      |
|----------------|----------------|
|    big.Int     |    uint128     |
|----------------|----------------|
|    big.Int     |    uint256     |
|----------------|----------------|
|     string     |    string      |
|----------------|----------------|
|common.Address* |    address     |
|----------------|----------------|
|    [32]byte    |    bytes32     |
|----------------|----------------|


*
common.Address is a specific go-ethereum type

```

## Workflow

### Testing
Project contains BDD tests based on [ginkgo](https://onsi.github.io/ginkgo/) tool.
Run tests:
```
make test
```
### Linting
Gollangci-lint used with `.golangci.yml` config obtained from [evmos](https://github.com/evmos/evmos) repository with little changes.

Run linter:
```
make lint
```

CI pipeline contains `build`, `lint` and `test` tasks.

[Conventional commits semantic](https://www.conventionalcommits.org/en/v1.0.0/) used in most of the commits.

## Further scope

Additional things, that may be done for the further development of this basic repository:

1. As every product created with time pressure this project have two common problems:
    - Tech debt consisting of non-reusable and not efficient code.
    - Lack of tests.
   So, these both points are a field for improvement.

2. Would be nice to implement a fully generic way of interaction with every function of smart contract.

3. Workflow and CI\CD improvements. For example, one of the first improvements may be packaging applications to docker containers.

4. Think about using more `evmos`/`ethermint` primetives and abstractions(for example for key pair generation or EVM RPC calls) as some of them may be usefull.
