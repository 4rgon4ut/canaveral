package cli

import (
	"fmt"
	"os"

	"github.com/canaveral/app"

	// ethermintclient "github.com/evmos/ethermint/client"
	"github.com/spf13/cobra"
)

type RootCmd struct {
	*cobra.Command
}

// Creates root cmd cobra command and register all subcommands
func NewRootCmd(a *app.App) *RootCmd {
	command := &cobra.Command{
		Use:   "canaveral",
		Short: "Compile and deploy smart contracts on EVM-compatible chains.",
	}
	command.AddCommand(
		compileCommand(a),
		deployCommand(a),
		callCommand(a),
	)
	return &RootCmd{command}
}

func (r *RootCmd) Run() {
	if err := r.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
