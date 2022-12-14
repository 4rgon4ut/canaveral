// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Test

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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_number\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"_word\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_obj\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFields\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"number\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"obj\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_number\",\"type\":\"uint16\"}],\"name\":\"setNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_obj\",\"type\":\"bytes32\"}],\"name\":\"setObj\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_word\",\"type\":\"string\"}],\"name\":\"setWord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"word\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620012d5380380620012d583398181016040528101906200003791906200036a565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600060146101000a81548161ffff021916908361ffff1602179055508260019081620000a5919062000646565b508160028190555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050506200072d565b6000604051905090565b600080fd5b600080fd5b600061ffff82169050919050565b62000125816200010c565b81146200013157600080fd5b50565b60008151905062000145816200011a565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620001a08262000155565b810181811067ffffffffffffffff82111715620001c257620001c162000166565b5b80604052505050565b6000620001d7620000f8565b9050620001e5828262000195565b919050565b600067ffffffffffffffff82111562000208576200020762000166565b5b620002138262000155565b9050602081019050919050565b60005b838110156200024057808201518184015260208101905062000223565b60008484015250505050565b6000620002636200025d84620001ea565b620001cb565b90508281526020810184848401111562000282576200028162000150565b5b6200028f84828562000220565b509392505050565b600082601f830112620002af57620002ae6200014b565b5b8151620002c18482602086016200024c565b91505092915050565b6000819050919050565b620002df81620002ca565b8114620002eb57600080fd5b50565b600081519050620002ff81620002d4565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620003328262000305565b9050919050565b620003448162000325565b81146200035057600080fd5b50565b600081519050620003648162000339565b92915050565b6000806000806080858703121562000387576200038662000102565b5b6000620003978782880162000134565b945050602085015167ffffffffffffffff811115620003bb57620003ba62000107565b5b620003c98782880162000297565b9350506040620003dc87828801620002ee565b9250506060620003ef8782880162000353565b91505092959194509250565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200044e57607f821691505b60208210810362000464576200046362000406565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620004ce7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200048f565b620004da86836200048f565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b600062000527620005216200051b84620004f2565b620004fc565b620004f2565b9050919050565b6000819050919050565b620005438362000506565b6200055b62000552826200052e565b8484546200049c565b825550505050565b600090565b6200057262000563565b6200057f81848462000538565b505050565b5b81811015620005a7576200059b60008262000568565b60018101905062000585565b5050565b601f821115620005f657620005c0816200046a565b620005cb846200047f565b81016020851015620005db578190505b620005f3620005ea856200047f565b83018262000584565b50505b505050565b600082821c905092915050565b60006200061b60001984600802620005fb565b1980831691505092915050565b600062000636838362000608565b9150826002028217905092915050565b6200065182620003fb565b67ffffffffffffffff8111156200066d576200066c62000166565b5b62000679825462000435565b62000686828285620005ab565b600060209050601f831160018114620006be5760008415620006a9578287015190505b620006b5858262000628565b86555062000725565b601f198416620006ce866200046a565b60005b82811015620006f857848901518255600182019150602085019450602081019050620006d1565b8683101562000718578489015162000714601f89168262000608565b8355505b6001600288020188555050505b505050505050565b610b98806200073d6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063ca60249211610066578063ca60249214610137578063cd048de614610158578063d1d80fdf14610174578063dfcfb27514610190578063e31f2d6e146101ac5761009e565b80632f64d386146100a3578063767800de146100c157806379d87e45146100df5780638381f58a146100fb5780638da5cb5b14610119575b600080fd5b6100ab6101ca565b6040516100b891906104a7565b60405180910390f35b6100c9610258565b6040516100d6919061050a565b60405180910390f35b6100f960048036038101906100f49190610573565b61027e565b005b61010361029e565b60405161011091906105af565b60405180910390f35b6101216102b2565b60405161012e919061050a565b60405180910390f35b61013f6102d6565b60405161014f94939291906105e3565b60405180910390f35b610172600480360381019061016d9190610764565b6103b0565b005b61018e600480360381019061018991906107d9565b6103c3565b005b6101aa60048036038101906101a59190610832565b610407565b005b6101b4610411565b6040516101c1919061085f565b60405180910390f35b600180546101d7906108a9565b80601f0160208091040260200160405190810160405280929190818152602001828054610203906108a9565b80156102505780601f1061022557610100808354040283529160200191610250565b820191906000526020600020905b81548152906001019060200180831161023357829003601f168201915b505050505081565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b80600060146101000a81548161ffff021916908361ffff16021790555050565b600060149054906101000a900461ffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006060600080600060149054906101000a900461ffff166001600254600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16828054610322906108a9565b80601f016020809104026020016040519081016040528092919081815260200182805461034e906108a9565b801561039b5780601f106103705761010080835404028352916020019161039b565b820191906000526020600020905b81548152906001019060200180831161037e57829003601f168201915b50505050509250935093509350935090919293565b80600190816103bf9190610a90565b5050565b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b8060028190555050565b60025481565b600081519050919050565b600082825260208201905092915050565b60005b83811015610451578082015181840152602081019050610436565b60008484015250505050565b6000601f19601f8301169050919050565b600061047982610417565b6104838185610422565b9350610493818560208601610433565b61049c8161045d565b840191505092915050565b600060208201905081810360008301526104c1818461046e565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104f4826104c9565b9050919050565b610504816104e9565b82525050565b600060208201905061051f60008301846104fb565b92915050565b6000604051905090565b600080fd5b600080fd5b600061ffff82169050919050565b61055081610539565b811461055b57600080fd5b50565b60008135905061056d81610547565b92915050565b6000602082840312156105895761058861052f565b5b60006105978482850161055e565b91505092915050565b6105a981610539565b82525050565b60006020820190506105c460008301846105a0565b92915050565b6000819050919050565b6105dd816105ca565b82525050565b60006080820190506105f860008301876105a0565b818103602083015261060a818661046e565b905061061960408301856105d4565b61062660608301846104fb565b95945050505050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6106718261045d565b810181811067ffffffffffffffff821117156106905761068f610639565b5b80604052505050565b60006106a3610525565b90506106af8282610668565b919050565b600067ffffffffffffffff8211156106cf576106ce610639565b5b6106d88261045d565b9050602081019050919050565b82818337600083830152505050565b6000610707610702846106b4565b610699565b90508281526020810184848401111561072357610722610634565b5b61072e8482856106e5565b509392505050565b600082601f83011261074b5761074a61062f565b5b813561075b8482602086016106f4565b91505092915050565b60006020828403121561077a5761077961052f565b5b600082013567ffffffffffffffff81111561079857610797610534565b5b6107a484828501610736565b91505092915050565b6107b6816104e9565b81146107c157600080fd5b50565b6000813590506107d3816107ad565b92915050565b6000602082840312156107ef576107ee61052f565b5b60006107fd848285016107c4565b91505092915050565b61080f816105ca565b811461081a57600080fd5b50565b60008135905061082c81610806565b92915050565b6000602082840312156108485761084761052f565b5b60006108568482850161081d565b91505092915050565b600060208201905061087460008301846105d4565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806108c157607f821691505b6020821081036108d4576108d361087a565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261093c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826108ff565b61094686836108ff565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b600061098d6109886109838461095e565b610968565b61095e565b9050919050565b6000819050919050565b6109a783610972565b6109bb6109b382610994565b84845461090c565b825550505050565b600090565b6109d06109c3565b6109db81848461099e565b505050565b5b818110156109ff576109f46000826109c8565b6001810190506109e1565b5050565b601f821115610a4457610a15816108da565b610a1e846108ef565b81016020851015610a2d578190505b610a41610a39856108ef565b8301826109e0565b50505b505050565b600082821c905092915050565b6000610a6760001984600802610a49565b1980831691505092915050565b6000610a808383610a56565b9150826002028217905092915050565b610a9982610417565b67ffffffffffffffff811115610ab257610ab1610639565b5b610abc82546108a9565b610ac7828285610a03565b600060209050601f831160018114610afa5760008415610ae8578287015190505b610af28582610a74565b865550610b5a565b601f198416610b08866108da565b60005b82811015610b3057848901518255600182019150602085019450602081019050610b0b565b86831015610b4d5784890151610b49601f891682610a56565b8355505b6001600288020188555050505b50505050505056fea264697066735822122036d5888c0f7198ca8e62ce1587853a1f03458914e2e2886fd8b08954592487e464736f6c63430008100033",
}

// TestABI is the input ABI used to generate the binding from.
// Deprecated: Use TestMetaData.ABI instead.
var TestABI = TestMetaData.ABI

// TestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestMetaData.Bin instead.
var TestBin = TestMetaData.Bin

// DeployTest deploys a new Ethereum contract, binding an instance of Test to it.
func DeployTest(auth *bind.TransactOpts, backend bind.ContractBackend, _number uint16, _word string, _obj [32]byte, _addr common.Address) (common.Address, *types.Transaction, *Test, error) {
	parsed, err := TestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestBin), backend, _number, _word, _obj, _addr)
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

// GetFields is a free data retrieval call binding the contract method 0xca602492.
//
// Solidity: function getFields() view returns(uint16, string, bytes32, address)
func (_Test *TestCaller) GetFields(opts *bind.CallOpts) (uint16, string, [32]byte, common.Address, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "getFields")

	if err != nil {
		return *new(uint16), *new(string), *new([32]byte), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return out0, out1, out2, out3, err

}

// GetFields is a free data retrieval call binding the contract method 0xca602492.
//
// Solidity: function getFields() view returns(uint16, string, bytes32, address)
func (_Test *TestSession) GetFields() (uint16, string, [32]byte, common.Address, error) {
	return _Test.Contract.GetFields(&_Test.CallOpts)
}

// GetFields is a free data retrieval call binding the contract method 0xca602492.
//
// Solidity: function getFields() view returns(uint16, string, bytes32, address)
func (_Test *TestCallerSession) GetFields() (uint16, string, [32]byte, common.Address, error) {
	return _Test.Contract.GetFields(&_Test.CallOpts)
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

// SetAddr is a paid mutator transaction binding the contract method 0xd1d80fdf.
//
// Solidity: function setAddr(address _addr) returns()
func (_Test *TestTransactor) SetAddr(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "setAddr", _addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd1d80fdf.
//
// Solidity: function setAddr(address _addr) returns()
func (_Test *TestSession) SetAddr(_addr common.Address) (*types.Transaction, error) {
	return _Test.Contract.SetAddr(&_Test.TransactOpts, _addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd1d80fdf.
//
// Solidity: function setAddr(address _addr) returns()
func (_Test *TestTransactorSession) SetAddr(_addr common.Address) (*types.Transaction, error) {
	return _Test.Contract.SetAddr(&_Test.TransactOpts, _addr)
}

// SetNumber is a paid mutator transaction binding the contract method 0x79d87e45.
//
// Solidity: function setNumber(uint16 _number) returns()
func (_Test *TestTransactor) SetNumber(opts *bind.TransactOpts, _number uint16) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "setNumber", _number)
}

// SetNumber is a paid mutator transaction binding the contract method 0x79d87e45.
//
// Solidity: function setNumber(uint16 _number) returns()
func (_Test *TestSession) SetNumber(_number uint16) (*types.Transaction, error) {
	return _Test.Contract.SetNumber(&_Test.TransactOpts, _number)
}

// SetNumber is a paid mutator transaction binding the contract method 0x79d87e45.
//
// Solidity: function setNumber(uint16 _number) returns()
func (_Test *TestTransactorSession) SetNumber(_number uint16) (*types.Transaction, error) {
	return _Test.Contract.SetNumber(&_Test.TransactOpts, _number)
}

// SetObj is a paid mutator transaction binding the contract method 0xdfcfb275.
//
// Solidity: function setObj(bytes32 _obj) returns()
func (_Test *TestTransactor) SetObj(opts *bind.TransactOpts, _obj [32]byte) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "setObj", _obj)
}

// SetObj is a paid mutator transaction binding the contract method 0xdfcfb275.
//
// Solidity: function setObj(bytes32 _obj) returns()
func (_Test *TestSession) SetObj(_obj [32]byte) (*types.Transaction, error) {
	return _Test.Contract.SetObj(&_Test.TransactOpts, _obj)
}

// SetObj is a paid mutator transaction binding the contract method 0xdfcfb275.
//
// Solidity: function setObj(bytes32 _obj) returns()
func (_Test *TestTransactorSession) SetObj(_obj [32]byte) (*types.Transaction, error) {
	return _Test.Contract.SetObj(&_Test.TransactOpts, _obj)
}

// SetWord is a paid mutator transaction binding the contract method 0xcd048de6.
//
// Solidity: function setWord(string _word) returns()
func (_Test *TestTransactor) SetWord(opts *bind.TransactOpts, _word string) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "setWord", _word)
}

// SetWord is a paid mutator transaction binding the contract method 0xcd048de6.
//
// Solidity: function setWord(string _word) returns()
func (_Test *TestSession) SetWord(_word string) (*types.Transaction, error) {
	return _Test.Contract.SetWord(&_Test.TransactOpts, _word)
}

// SetWord is a paid mutator transaction binding the contract method 0xcd048de6.
//
// Solidity: function setWord(string _word) returns()
func (_Test *TestTransactorSession) SetWord(_word string) (*types.Transaction, error) {
	return _Test.Contract.SetWord(&_Test.TransactOpts, _word)
}
