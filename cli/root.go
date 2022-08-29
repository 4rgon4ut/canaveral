package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/canaveral/app"

	// ethermintclient "github.com/evmos/ethermint/client"
	"github.com/spf13/cobra"
)

//
type RootCmd struct {
	*cobra.Command
}

// Creates root cmd cobra command and register all subcommands
func NewRootCmd(a *app.App) *RootCmd {
	command := &cobra.Command{
		Use:   "cn",
		Short: "Compile and deploy smart contracts on EVM-compatible chains.",
	}
	command.AddCommand(
		compileCommand(a),
		deployCommand(a),

		// erc20 example specific funcs
		deployExampleERC20Command(a),
		balanceCommand(a),
		transferCommand(a),
		mintCommand(a),
		burnCommand(a),

		// example a.k.a. DEMO command
		exampleCommand(a),
	)
	return &RootCmd{command}
}

func (r *RootCmd) Run() {
	if err := r.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

//nolint:all
// Example command for demonstration of compilation, deploy and interaction with ExampleERC20 token contract
func exampleCommand(a *app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "example",
		Short: "Starts example run of ERC20 interaction",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("\n\n                  WELCOME TO CANAVERAL\n\n\n")
			time.Sleep(time.Second * 4)
			fmt.Println("We are going to compile our ExampleERC20 token...\n\n")
			time.Sleep(time.Second * 4)
			if err := a.Compile("ExampleERC20.sol"); err != nil {
				return fmt.Errorf("compile command error: %w", err)
			}
			fmt.Println("______________________________________________________________\n\n")
			time.Sleep(time.Second * 13)
			fmt.Println("Now we need to launch our rocket...(deploy compiled contract)\n\n")
			time.Sleep(time.Second * 5)
			fmt.Println("3...\n")
			time.Sleep(time.Second * 1)
			fmt.Println("2...\n")
			time.Sleep(time.Second * 1)
			fmt.Println("1...\n")
			time.Sleep(time.Second * 2)
			if err := a.DeployExampleERC20(); err != nil {
				fmt.Println("Rocket crushed... :(")
				return fmt.Errorf("compile command error: %w", err)
			}
			fmt.Println("______________________________________________________________\n\n")
			time.Sleep(time.Second * 9)
			fmt.Println("Now we can interact with our ERC20 token contract \n")
			time.Sleep(time.Second * 6)
			fmt.Println("Firstly we need to mint some tokens: \n")
			time.Sleep(time.Second * 6)
			if err := a.Mint(a.EVMClient.Account.Address.Hex(), "100000"); err != nil {
				return fmt.Errorf("mint command error: %w", err)
			}
			fmt.Println("______________________________________________________________\n\n")
			time.Sleep(time.Second * 6)
			fmt.Println("Now get our balance: \n")
			time.Sleep(time.Second * 6)
			if err := a.BalanceOf(a.EVMClient.Account.Address.Hex()); err != nil {
				return fmt.Errorf("balanceOf command error: %w", err)
			}
			fmt.Println("______________________________________________________________\n\n")
			time.Sleep(time.Second * 6)
			fmt.Println("Transfer our tokens: \n")
			time.Sleep(time.Second * 6)
			if err := a.Transfer("0xb08849e93CD2F2eAb3eF0cBF90faEF3ba3eea882", "1000"); err != nil {
				return fmt.Errorf("transfer command error: %w", err)
			}
			fmt.Println("______________________________________________________________\n\n")
			time.Sleep(time.Second * 6)
			fmt.Println("And finally burn: \n")
			time.Sleep(time.Second * 6)
			if err := a.Burn("1000"); err != nil {
				return fmt.Errorf("burn command error: %w", err)
			}
			fmt.Println("______________________________________________________________\n\n")
			time.Sleep(time.Second * 6)
			fmt.Println("Thank you! \nNow you can interact with deployed contract by yourself with cli commands.")
			time.Sleep(time.Second * 6)
			return nil
		},
	}
}
