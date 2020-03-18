package cmd

import (
	"github.com/m-porter/arc/lib/config"
	"github.com/m-porter/arc/lib/sync"
	"github.com/m-porter/arc/lib/util"
	"github.com/spf13/cobra"
)

func newSyncCmd() *cobra.Command {
	var forceSync bool

	syncCmd := &cobra.Command{
		Use:   "sync",
		Short: "Syncs all defined services for the active project",
		Long:  "Syncs all defined services for the active project. By default, dirty branches and the current working directory service will not be sync'd.",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.LoadArcConfig()
			err := sync.ProjectByName(cfg.CurrentProject, forceSync)
			if err != nil {
				util.Fatalf("%v", err)
			}
		},
	}

	syncCmd.Flags().BoolVar(&forceSync, "force", false, "force sync dirty branches")

	return syncCmd
}
