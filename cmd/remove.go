package cmd

import (
	"github.com/m-porter/arc/lib/resource"
	"github.com/m-porter/arc/lib/util"
	"github.com/spf13/cobra"
)

func newRemoveCmd() *cobra.Command {
	var resourceName string
	var projectName string

	removeCmd := &cobra.Command{
		Use:     "remove (project|service) --name=NAME",
		Aliases: []string{"rm"},
		Short:   "Remove an arc resource",
		Example: `arc remove project --name=spd-integration
arc remove service --name=chamber --project=spd-integration`,
		ValidArgs: []string{"project", "service"},
		Args:      cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			resourceType := args[0]

			util.EnforceFlag(resourceName, "", "--name flag required")

			switch resourceType {
			case "project":
				resource.RemoveProject(resourceName)
			case "service":
				util.EnforceFlag(projectName, "", "--project flag required")
				resource.RemoveService(resourceName, projectName)
			}
		},
	}

	removeCmd.Flags().StringVar(&resourceName, "name", "", "resource name")
	removeCmd.Flags().StringVar(&projectName, "project", "", "project name")

	return removeCmd
}
