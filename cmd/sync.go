package cmd

import (
	"github.com/spf13/cobra"
)

func newSyncCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "sync",
		Short: "Syncs all defined services for the active project",
		Run: func(cmd *cobra.Command, args []string) {
			//
		},
	}
}
