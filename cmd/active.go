package cmd

import (
	"github.com/spf13/cobra"
)

func newActiveCmd() *cobra.Command {
	var projectName string

	activeCmd := &cobra.Command{
		Use:   "active --project=PROJECT",
		Short: "Show or set the active arc project",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			//
		},
	}

	activeCmd.Flags().StringVarP(&projectName, "project", "P", "", "project name")

	return activeCmd
}
