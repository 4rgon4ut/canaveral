package utils

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Takes path to the *.abi file, read it and returns parsed ABI interface object
func GetABIObject(abiPath string) (*abi.ABI, error) {
	b, err := os.ReadFile(abiPath)
	if err != nil {
		return nil, err
	}
	evmABI, err := abi.JSON(strings.NewReader(string(b)))
	if err != nil {
		return nil, err
	}
	return &evmABI, nil
}

// Takes path to the *.bin file, read file and returns byte array representation
func GetBytecode(binPath string) ([]byte, error) {
	return os.ReadFile(binPath)
}

// Searches for method with provided name in contract ABI
// TODO: overloaded methods collision
// TODO: not efficient
func GetMethodByName(contractABI abi.ABI, name string) (abi.Method, bool) {
	for _, m := range contractABI.Methods {
		if m.Name == name {
			return m, true
		}
	}
	return abi.Method{}, false
}

// Cast types of args([ ]string) to types of provided [ ]Arguments.
//
// abiArgs[ Argument{Type: type1}, Argument{Type: type2} ] --> args[0].(type1), args[1].(type2), ...
//
// Function supports only several types(any <uintX>, <string>), plus ETH <common.Address> and <[32]byte>
// and is not safe to buffer overflow
//
// TODO: add checks for resulting types overflow
// TODO: spread supported types number
func CastInputs(abiArgs abi.Arguments, args []string) ([]interface{}, error) {
	argsLen := len(args)
	abiArgsLen := len(abiArgs)

	if abiArgsLen != argsLen {
		return nil, fmt.Errorf("ivalid arguments number: expect %d got %d", abiArgsLen, argsLen)
	}

	assertedArgs := make([]interface{}, argsLen)
	for i, arg := range abiArgs {
		stringArg := args[i]
		// string representation of solidity type to find a match
		switch arg.Type.String() {
		case "uint8":
			u64, err := strconv.ParseUint(stringArg, 10, 8)
			if err != nil {
				return nil, err
			}
			assertedArgs[i] = uint8(u64)
		case "uint16":
			u64, err := strconv.ParseUint(stringArg, 10, 16)
			if err != nil {
				return nil, err
			}
			assertedArgs[i] = uint16(u64)
		case "uint32":
			u64, err := strconv.ParseUint(stringArg, 10, 32)
			if err != nil {
				return nil, err
			}
			assertedArgs[i] = uint32(u64)
		case "uint64":
			u64, err := strconv.ParseUint(stringArg, 10, 64)
			if err != nil {
				return nil, err
			}
			assertedArgs[i] = u64
		case "uint256", "uint128":
			n := new(big.Int)
			n, ok := n.SetString(stringArg, 10)
			if !ok {
				return nil, fmt.Errorf("can't convert %s to big.Int", stringArg)
			}
			assertedArgs[i] = n
		case "string":
			assertedArgs[i] = stringArg
		case "address":
			if !common.IsHexAddress(stringArg) {
				return nil, fmt.Errorf("invalid address provided: %x", stringArg)
			}
			assertedArgs[i] = common.HexToAddress(stringArg)
		case "bytes32":
			var param [32]uint8
			copy(param[:], stringArg)
			assertedArgs[i] = param

		default:
			return nil, fmt.Errorf("unsupported input type")
		}
	}
	return assertedArgs, nil
}
