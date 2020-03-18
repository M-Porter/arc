package cmd

import (
	"github.com/m-porter/arc/lib/resource"
	"github.com/m-porter/arc/lib/util"
	"github.com/spf13/cobra"
)

func newNewCmd() *cobra.Command {
	var resourceName string
	var projectName string
	var servicePath string
	var serviceBranch string
	var setActive bool

	newCmd := &cobra.Command{
		Use:   "new (project|service) --name=NAME",
		Short: "Create a new arc resource",
		Example: `new project --name="spd-integration"
new service --name="chamber" --path="/absolute/local/chamber" --branch="spd-integration" --project="spd-integration"`,
		ValidArgs: []string{"project", "service"},
		Args:      cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			resourceType := args[0]

			util.EnforceFlag(resourceName, "", "--name flag required")

			switch resourceType {
			case "project":
				resource.CreateProject(resourceName, setActive)
			case "service":
				util.EnforceFlag(servicePath, "", "--path flag required")
				resource.CreateService(resourceName, projectName, servicePath, serviceBranch)
			}
		},
	}

	newCmd.Flags().StringVar(&resourceName, "name", "", "resource name")
	newCmd.Flags().BoolVar(&setActive, "active", false, "project only - sets the active project to this")
	newCmd.Flags().StringVar(&projectName, "project", "", "service only - the project to assign this new service to")
	newCmd.Flags().StringVar(&servicePath, "path", "", "service only - the absolute path to the local git repository")
	newCmd.Flags().StringVar(&serviceBranch, "branch", "", "service only - the git branch to use")

	return newCmd
}
