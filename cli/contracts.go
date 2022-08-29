package cli

import (
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
			if err := a.Deploy(args[0], args[1:]); err != nil {
				return err
			}
			return nil
		},
	}
}
