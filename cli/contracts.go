package cli

import (
	"fmt"

	"github.com/canaveral/app"
	"github.com/spf13/cobra"
)

func compileCommand(a *app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "compile [contract_name]",
		Short: "Compile *.sol file into *.abi && *.bin",
		Long:  `This command receives a contract name, compiles ABI and binary files and creates bind with *.go file`,
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := a.Compile(args[0]); err != nil {
				return err
			}
			return nil
		},
	}
}

func deployCommand(a *app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "deploy [contract_name] [constructor input]",
		Short: "Deploy contract to the network",
		Long:  `This command receives a contract name and deploys it to the blockchain`,
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, tx, err := a.Deploy(args[0], args[1:])
			if err != nil {
				return err
			}
			fmt.Printf(
				"%s seccessfully deployed. \nTX hash: %s \nContract address: %s\n\n",
				args[0],
				tx,
				addr,
			)
			return nil
		},
	}
}

func callCommand(a *app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "call [contract_name] [method] [input]",
		Short: "",
		Long:  ``,
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := a.Call(args[0], args[1], args[2:]); err != nil {
				fmt.Println(err)
				return err
			}
			return nil
		},
	}
}
