package cmd

import (
	"github.com/spf13/cobra"
)

func newSyncCmd() *cobra.Command {
	var forceSync bool

	syncCmd := &cobra.Command{
		Use:   "sync",
		Short: "Syncs all defined services for the active project",
		Long:  "Syncs all defined services for the active project. By default, dirty branches and the current working directory service will not be sync'd.",
		Run: func(cmd *cobra.Command, args []string) {
			//
		},
	}

	syncCmd.Flags().BoolVarP(&forceSync, "force", "f", false, "force sync dirty branches")

	return syncCmd
}
