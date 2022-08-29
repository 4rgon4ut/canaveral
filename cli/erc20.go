package cli

import (
	"github.com/canaveral/app"
	"github.com/spf13/cobra"
)

func deployExampleERC20Command(a *app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "deploy-example",
		Short: "Deploys example token contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := a.DeployExampleERC20(); err != nil {
				return err
			}
			return nil
		},
	}
}

func balanceCommand(a *app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "balance [address]",
		Short: "Query example token contract to get address balance",
		Long:  `This command receives an account address, performs view function balanceOf call to get balance`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := a.BalanceOf(args[0]); err != nil {
				return err
			}
			return nil
		},
	}
}

func transferCommand(a *app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "transfer [address] [amount]",
		Short: "Transfer amount of tokens to provided address",
		Long:  `This command receives a recipient address and amount of Example contract tokens to transfer`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := a.Transfer(args[0], args[1]); err != nil {
				return err
			}
			return nil
		},
	}
}

func mintCommand(a *app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "mint [address] [amount]",
		Short: "Mint provided amount of ExampleERC20 tokens",
		Long:  `This command receives an amount of tokens and recipient address to perform mint`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := a.Mint(args[0], args[1]); err != nil {
				return err
			}
			return nil
		},
	}
}

func burnCommand(a *app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "burn [amount]",
		Short: "Burn provided amount of ExampleERC20 tokens",
		Long:  `This command receives an amount of tokens to call burn function on ExampleERC20 contract`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := a.Burn(args[0]); err != nil {
				return err
			}
			return nil
		},
	}
}
