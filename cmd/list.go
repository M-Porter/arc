package cmd

import "github.com/spf13/cobra"

func newListCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Lists all projects and their services",
		Run: func(cmd *cobra.Command, args []string) {
			//
		},
	}
}
