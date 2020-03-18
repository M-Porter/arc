package cmd

import (
	"fmt"
	"github.com/m-porter/arc/lib/config"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Lists all projects and their services",
		Run: func(cmd *cobra.Command, args []string) {
			projects := config.GetArcConfig().Projects

			if len(projects) == 0 {
				fmt.Println("no defined projects")
				return
			}

			config.Println(projects)
		},
	}
}
