package cmd

import (
	"github.com/m-porter/arc/lib/resource"
	"github.com/m-porter/arc/lib/util"
	"github.com/spf13/cobra"
)

func newNewCommand() *cobra.Command {
	options := resource.CreateResourceOptions{}

	newCmd := &cobra.Command{
		Use:   "new (project|service) --name=NAME",
		Short: "Create a new arc resource",
		Example: `new project --name="spd-integration"
new service --name="chamber" --path="/absolute/local/chamber" --branch="spd-integration" --project="spd-integration"`,
		ValidArgs: []string{"project", "service"},
		Args:      cobra.ExactValidArgs(1),
		Run:       newCommandHandler(&options),
	}

	newCmd.Flags().StringVarP(&options.ResourceName, "name", "n", "", "resource name")
	newCmd.Flags().StringVarP(&options.ProjectName, "project", "P", "", "service only - the project to assign this new service to")
	newCmd.Flags().StringVarP(&options.ServicePath, "path", "p", "", "service only - the absolute path to the local git repository")
	newCmd.Flags().StringVarP(&options.ServiceBranch, "branch", "b", "", "service only - the git branch to use")

	return newCmd
}

func newCommandHandler(options *resource.CreateResourceOptions) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		resourceType := args[0]

		util.EnforceFlag(options.ResourceName, "", "--name flag required")

		switch resourceType {
		case "project":
			resource.CreateProject(options)
		case "service":
			util.EnforceFlag(options.ServicePath, "", "--path flag required")
			resource.CreateService(options)
		}
	}
}
