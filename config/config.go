package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	//
	RPCAddr    string `mapstructure:"RPC_ADDR"`
	RPCPort    string `mapstructure:"RPC_PORT"`
	PrivateKey string `mapstructure:"PRIVATE_KEY"`

	ERC20ExampleAddress string `mapstructure:"ERC20_EXAMPLE_ADDRESS"`
	//
	BinDir       string `mapstructure:"BIN_DIR"`
	ABIDir       string `mapstructure:"ABI_DIR"`
	BindsDir     string `mapstructure:"BINDS_DIR"`
	ContractsDir string `mapstructure:"CONTRACTS_DIR"`

	LogLevel string `mapstrucutre:"LOG_LEVEL"`
}

func AddExampleContractAddr(addr string) error {
	viper.Set("ERC20_EXAMPLE_ADDRESS", addr)
	if err := viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}

// Loads config data and unmarshal it to Config struct.
func Init(path string, name string, configType string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)       // Register config file name (no extension)
	viper.SetConfigType(configType) // Look for specific type
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("config loading error: %w", err)
	}
	c := &Config{}
	if err := viper.Unmarshal(c); err != nil {
		return nil, fmt.Errorf("unmarshaling config error: %w", err)
	}
	return c, nil
}
