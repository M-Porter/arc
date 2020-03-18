package cmd

import (
	"fmt"
	"github.com/m-porter/arc/lib/config"
	"github.com/m-porter/arc/lib/util"
	"github.com/spf13/cobra"
)

func newActiveCmd() *cobra.Command {
	var projectName string

	activeCmd := &cobra.Command{
		Use:   "active --project=PROJECT",
		Short: "Show or set the active arc project",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.LoadArcConfig()

			if projectName == "" {
				fmt.Println(cfg.CurrentProject)
				return
			}

			util.Printlnf("activating %v", projectName)
			cfg.CurrentProject = projectName
			config.WriteArcConfig(cfg)
			util.Printlnf(`run "arc sync" to update all services`)
		},
	}

	activeCmd.Flags().StringVarP(&projectName, "project", "P", "", "project name")

	return activeCmd
}
