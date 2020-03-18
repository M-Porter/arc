package cmd

import (
	"fmt"
	"github.com/m-porter/arc/lib/config"
	"github.com/m-porter/arc/lib/sync"
	"github.com/m-porter/arc/lib/util"
	"github.com/spf13/cobra"
)

func newActiveCmd() *cobra.Command {
	var projectName string
	var autoSync bool

	activeCmd := &cobra.Command{
		Use:   "active --project=PROJECT",
		Short: "Show or set the active arc project",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.LoadArcConfig()

			if projectName == "" {
				fmt.Println(cfg.ActiveProject)
				return
			}

			util.Printlnf("activating %v", projectName)
			cfg.ActiveProject = projectName
			config.WriteArcConfig(cfg)

			if autoSync {
				err := sync.ProjectByName(cfg.ActiveProject, false)
				if err != nil {
					util.Fatalf("%v", err)
				}
			} else {
				util.Printlnf(`run "arc sync" to update all services`)
			}
		},
	}

	activeCmd.Flags().StringVar(&projectName, "project", "", "project name")
	activeCmd.Flags().BoolVar(&autoSync, "sync", false, "sync project after activating; only works if --project option provided")

	return activeCmd
}
