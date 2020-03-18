package cmd

import (
	"github.com/m-porter/arc/lib/config"
	"github.com/m-porter/arc/lib/utils"
	"github.com/spf13/cobra"
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
	config.EnsureArcConfig()

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
		utils.Fatalf("error starting arc: %v", err)
	}
}
