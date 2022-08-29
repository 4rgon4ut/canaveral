package config_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/canaveral/config"
)

var _ = Describe("Config", func() {
	var predefined *config.Config

	BeforeEach(func() {
		predefined = &config.Config{
			RPCAddr:    "http://0.0.0.0",
			RPCPort:    "8545",
			PrivateKey: "41221710A88C7284D5F5BDE53311CDFA22B20F2724F3F4DA60FE3F567669DF56",

			BinDir:       "artifacts/bin",
			ABIDir:       "artifacts/abi",
			BindsDir:     "bindings",
			ContractsDir: "contracts",
			LogLevel:     "",
		}

	})
	Context("with right config path, name and type", func() {
		It("predifined and loaded configs should be equal", func() {
			loaded, _ := config.Init(".", "config", "toml")
			Expect(loaded.RPCAddr).To(Equal(predefined.RPCAddr))
			Expect(loaded.RPCPort).To(Equal(predefined.RPCPort))
			Expect(loaded.PrivateKey).To(Equal(predefined.PrivateKey))
			Expect(loaded.BinDir).To(Equal(predefined.BinDir))
			Expect(loaded.ABIDir).To(Equal(predefined.ABIDir))
			Expect(loaded.BindsDir).To(Equal(predefined.BindsDir))
			Expect(loaded.ContractsDir).To(Equal(predefined.ContractsDir))
			Expect(loaded.LogLevel).To(Equal(predefined.LogLevel))
		})
	})
	Context("with wrong config path, name or type", func() {
		It("Init func should return error", func() {
			_, err := config.Init("gfegeefw", "ewffew", "ewfew")
			Expect(err).Error()
			fmt.Println(err)
		})
		It("loaded config should be nil", func() {
			loadedNil, _ := config.Init("wdfdwfdwfew", "fwdfwf", "fwfewf")
			Expect(loadedNil).To(BeNil())
		})
	})

})
