package cmd

import (
	"github.com/spf13/cobra"
)

func newNewCommand() *cobra.Command {
	var resourceName string
	var projectName string
	var servicePath string
	var serviceBranch string

	newCmd := &cobra.Command{
		Use:   "new (project|service) --name=NAME",
		Short: "Create a new arc resource",
		Example: `new project --name="spd-integration"
new service --name="chamber" --path="/absolute/local/chamber" --branch="spd-integration" --project="spd-integration"`,
		ValidArgs: []string{"project", "service"},
		Args:      cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			//
		},
	}

	newCmd.Flags().StringVarP(&resourceName, "name", "n", "", "resource name")
	newCmd.Flags().StringVarP(&projectName, "project", "P", "", "service only - the project to assign this new service to")
	newCmd.Flags().StringVarP(&servicePath, "path", "p", "", "service only - the absolute path to the local git repository")
	newCmd.Flags().StringVarP(&serviceBranch, "branch", "b", "", "service only - the git branch to use")

	return newCmd
}
