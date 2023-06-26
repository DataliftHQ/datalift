package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type migrateCmd struct {
	cmd  *cobra.Command
	opts migrateOpts
}

type migrateOpts struct {
}

func newMigrateCmd() *migrateCmd {
	root := &migrateCmd{}
	cmd := &cobra.Command{
		Use:               "migrate",
		Aliases:           []string{"m"},
		Short:             "Generates GoReleaser's command line docs",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Migrate")
			return nil
		},
	}
	root.cmd = cmd
	return root
}
