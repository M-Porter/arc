package cmd

import (
	"github.com/spf13/cobra"
)

func newRemoveCmd() *cobra.Command {
	var resourceName string
	var projectName string

	removeCmd := &cobra.Command{
		Use:     "remove (project|service) --name=NAME",
		Aliases: []string{"rm"},
		Short:   "Remove an arc resource",
		Example: `remove project --name=spd-integration
remove service --name=chamber --project=spd-integration`,
		ValidArgs: []string{"project", "service"},
		Args:      cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			//
		},
	}

	removeCmd.Flags().StringVarP(&resourceName, "name", "n", "", "resource name")
	removeCmd.Flags().StringVarP(&projectName, "project", "P", "", "project name")

	return removeCmd
}
