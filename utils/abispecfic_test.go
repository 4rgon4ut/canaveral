package utils_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/canaveral/utils"
)

var _ = Describe("Abispecfic", func() {

	var testABI *abi.ABI

	BeforeEach(func() {
		var TestMetaData = &bind.MetaData{
			ABI: "[{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_number\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"_word\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_obj\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFields\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"number\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"obj\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_number\",\"type\":\"uint16\"}],\"name\":\"setNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_obj\",\"type\":\"bytes32\"}],\"name\":\"setObj\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_word\",\"type\":\"string\"}],\"name\":\"setWord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"word\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
		}
		testABI, _ = TestMetaData.GetAbi()

	})

	Context("GetABIObject", func() {
		abiPath := "../artifacts/abi/Test.abi"
		It("should load and return abi type object without error", func() {
			object, err := utils.GetABIObject(abiPath)
			expectedType := new(abi.ABI)
			Expect(err).To(BeNil())
			Expect(object).To(BeAssignableToTypeOf(expectedType))
		})
		It("should return error on invalid .abi file path", func() {
			object, err := utils.GetABIObject("dwefwefwefewwfe")
			Expect(err).Error()
			Expect(err).ToNot(BeNil())
			Expect(object).To(BeNil())
		})
	})

	Context("GetBytecode", func() {
		binPath := "../artifacts/bin/Test.bin"
		It("should read .bin file and return bytes array representation", func() {
			object, err := utils.GetBytecode(binPath)
			expectedType := []byte{}
			Expect(err).To(BeNil())
			Expect(object).To(BeAssignableToTypeOf(expectedType))
		})
		It("should return error on invalid .bin file path", func() {
			object, err := utils.GetABIObject("dwefwefwefewwfe")
			Expect(err).Error()
			Expect(err).ToNot(BeNil())
			Expect(object).To(BeNil())
		})
	})

	Context("GetMethodByName", func() {
		It("should find and return abi.Method with 'true' flag", func() {
			methodName := "setNumber"
			method, exist := utils.GetMethodByName(*testABI, methodName)
			Expect(exist).To(BeTrue())
			Expect(method.Name).To(Equal(methodName))
		})
		It("should return empty abi.Method with 'false' flag if no such method", func() {
			methodName := "adsasdasd"
			method, exist := utils.GetMethodByName(*testABI, methodName)
			Expect(exist).To(BeFalse())
			Expect(method.Name).To(Equal(""))
		})
	})

	Context("CastInputs", func() {
		var b [32]byte
		// slice with correct types to match after cast
		correctlyTypedSlice := []interface{}{uint16(666), "Hola!", b, common.HexToAddress("0xA0FA3CaCDc397F49C4D19221Bcef495b283D2422")}
		// user input
		argsToCast := []string{"666", "Hola!", "string to bytes", "0xA0FA3CaCDc397F49C4D19221Bcef495b283D2422"}

		It("should cast slice of strings to provided function type input params types", func() {
			contractArgs := testABI.Constructor.Inputs
			inputs, err := utils.CastInputs(contractArgs, argsToCast)
			Expect(err).To(BeNil())
			Expect(len(inputs)).To(Equal(len(argsToCast)))
			// types matching
			for i, arg := range inputs {
				Expect(arg).To(BeAssignableToTypeOf(correctlyTypedSlice[i]))
			}
		})

		It("should return error if unable to cast", func() {
			// make first arg of expected type uint16 unconvertable(string)
			contractArgs := testABI.Constructor.Inputs
			argsToCast[0] = "csadcds"
			inputs, err := utils.CastInputs(contractArgs, argsToCast)
			Expect(err).ToNot(BeNil())
			Expect(inputs).To(BeNil())
		})
	})
})
