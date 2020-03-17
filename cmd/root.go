package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "arc",
		Short: "Bring electricity to your development experience",
		Long: `█████╗ ██████╗  ██████╗
██╔══██╗██╔══██╗██╔════╝
███████║██████╔╝██║
██╔══██║██╔══██╗██║
██║  ██║██║  ██║╚██████╗
╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝

Bring electricity to your development experience`,
	}
}

func Execute() {
	rootCmd := newRootCmd()

	rootCmd.AddCommand(
		newNewCommand(),
		newRemoveCmd(),
		newActiveCmd(),
		newListCmd(),
		newSyncCmd(),
	)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
