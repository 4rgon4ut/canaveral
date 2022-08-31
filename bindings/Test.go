// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TestMetaData contains all meta data concerning the Test contract.
var TestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_number\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"number\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"obj\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"word\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610504380380610504833981810160405281019061003291906100d4565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548161ffff021916908361ffff16021790555050610101565b600080fd5b600061ffff82169050919050565b6100b18161009a565b81146100bc57600080fd5b50565b6000815190506100ce816100a8565b92915050565b6000602082840312156100ea576100e9610095565b5b60006100f8848285016100bf565b91505092915050565b6103f4806101106000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80632f64d3861461005c578063767800de1461007a5780638381f58a146100985780638da5cb5b146100b6578063e31f2d6e146100d4575b600080fd5b6100646100f2565b6040516100719190610274565b60405180910390f35b610082610180565b60405161008f91906102d7565b60405180910390f35b6100a06101a6565b6040516100ad919061030f565b60405180910390f35b6100be6101ba565b6040516100cb91906102d7565b60405180910390f35b6100dc6101de565b6040516100e99190610343565b60405180910390f35b600180546100ff9061038d565b80601f016020809104026020016040519081016040528092919081815260200182805461012b9061038d565b80156101785780601f1061014d57610100808354040283529160200191610178565b820191906000526020600020905b81548152906001019060200180831161015b57829003601f168201915b505050505081565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900461ffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b600081519050919050565b600082825260208201905092915050565b60005b8381101561021e578082015181840152602081019050610203565b60008484015250505050565b6000601f19601f8301169050919050565b6000610246826101e4565b61025081856101ef565b9350610260818560208601610200565b6102698161022a565b840191505092915050565b6000602082019050818103600083015261028e818461023b565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006102c182610296565b9050919050565b6102d1816102b6565b82525050565b60006020820190506102ec60008301846102c8565b92915050565b600061ffff82169050919050565b610309816102f2565b82525050565b60006020820190506103246000830184610300565b92915050565b6000819050919050565b61033d8161032a565b82525050565b60006020820190506103586000830184610334565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806103a557607f821691505b6020821081036103b8576103b761035e565b5b5091905056fea26469706673582212209de48291dd0e0e94791dae844012f7170f4d96dae18f82aa3165d4fc4dae244764736f6c63430008100033",
}

// TestABI is the input ABI used to generate the binding from.
// Deprecated: Use TestMetaData.ABI instead.
var TestABI = TestMetaData.ABI

// TestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestMetaData.Bin instead.
var TestBin = TestMetaData.Bin

// (d Deployer) DeployTest deploys a new Ethereum contract, binding an instance of Test to it.
func (d Deployer) DeployTest(auth *bind.TransactOpts, backend bind.ContractBackend, _number uint16) (common.Address, *types.Transaction, *Test, error) {
	parsed, err := TestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestBin), backend, _number)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}, TestFilterer: TestFilterer{contract: contract}}, nil
}

// Test is an auto generated Go binding around an Ethereum contract.
type Test struct {
	TestCaller     // Read-only binding to the contract
	TestTransactor // Write-only binding to the contract
	TestFilterer   // Log filterer for contract events
}

// TestCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestSession struct {
	Contract     *Test             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestCallerSession struct {
	Contract *TestCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestTransactorSession struct {
	Contract     *TestTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestRaw struct {
	Contract *Test // Generic contract binding to access the raw methods on
}

// TestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestCallerRaw struct {
	Contract *TestCaller // Generic read-only contract binding to access the raw methods on
}

// TestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestTransactorRaw struct {
	Contract *TestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTest creates a new instance of Test, bound to a specific deployed contract.
func NewTest(address common.Address, backend bind.ContractBackend) (*Test, error) {
	contract, err := bindTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}, TestFilterer: TestFilterer{contract: contract}}, nil
}

// NewTestCaller creates a new read-only instance of Test, bound to a specific deployed contract.
func NewTestCaller(address common.Address, caller bind.ContractCaller) (*TestCaller, error) {
	contract, err := bindTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestCaller{contract: contract}, nil
}

// NewTestTransactor creates a new write-only instance of Test, bound to a specific deployed contract.
func NewTestTransactor(address common.Address, transactor bind.ContractTransactor) (*TestTransactor, error) {
	contract, err := bindTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestTransactor{contract: contract}, nil
}

// NewTestFilterer creates a new log filterer instance of Test, bound to a specific deployed contract.
func NewTestFilterer(address common.Address, filterer bind.ContractFilterer) (*TestFilterer, error) {
	contract, err := bindTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestFilterer{contract: contract}, nil
}

// bindTest binds a generic wrapper to an already deployed contract.
func bindTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test.Contract.TestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.contract.Transact(opts, method, params...)
}

// Addr is a free data retrieval call binding the contract method 0x767800de.
//
// Solidity: function addr() view returns(address)
func (_Test *TestCaller) Addr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "addr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0x767800de.
//
// Solidity: function addr() view returns(address)
func (_Test *TestSession) Addr() (common.Address, error) {
	return _Test.Contract.Addr(&_Test.CallOpts)
}

// Addr is a free data retrieval call binding the contract method 0x767800de.
//
// Solidity: function addr() view returns(address)
func (_Test *TestCallerSession) Addr() (common.Address, error) {
	return _Test.Contract.Addr(&_Test.CallOpts)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint16)
func (_Test *TestCaller) Number(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "number")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint16)
func (_Test *TestSession) Number() (uint16, error) {
	return _Test.Contract.Number(&_Test.CallOpts)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint16)
func (_Test *TestCallerSession) Number() (uint16, error) {
	return _Test.Contract.Number(&_Test.CallOpts)
}

// Obj is a free data retrieval call binding the contract method 0xe31f2d6e.
//
// Solidity: function obj() view returns(bytes32)
func (_Test *TestCaller) Obj(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "obj")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Obj is a free data retrieval call binding the contract method 0xe31f2d6e.
//
// Solidity: function obj() view returns(bytes32)
func (_Test *TestSession) Obj() ([32]byte, error) {
	return _Test.Contract.Obj(&_Test.CallOpts)
}

// Obj is a free data retrieval call binding the contract method 0xe31f2d6e.
//
// Solidity: function obj() view returns(bytes32)
func (_Test *TestCallerSession) Obj() ([32]byte, error) {
	return _Test.Contract.Obj(&_Test.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Test *TestCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Test *TestSession) Owner() (common.Address, error) {
	return _Test.Contract.Owner(&_Test.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Test *TestCallerSession) Owner() (common.Address, error) {
	return _Test.Contract.Owner(&_Test.CallOpts)
}

// Word is a free data retrieval call binding the contract method 0x2f64d386.
//
// Solidity: function word() view returns(string)
func (_Test *TestCaller) Word(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "word")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Word is a free data retrieval call binding the contract method 0x2f64d386.
//
// Solidity: function word() view returns(string)
func (_Test *TestSession) Word() (string, error) {
	return _Test.Contract.Word(&_Test.CallOpts)
}

// Word is a free data retrieval call binding the contract method 0x2f64d386.
//
// Solidity: function word() view returns(string)
func (_Test *TestCallerSession) Word() (string, error) {
	return _Test.Contract.Word(&_Test.CallOpts)
}
