package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type workerCmd struct {
	cmd  *cobra.Command
	opts workerOpts
}

type workerOpts struct {
}

func newWorkerCmd() *workerCmd {
	root := &workerCmd{}
	cmd := &cobra.Command{
		Use:               "worker",
		Aliases:           []string{"w"},
		Short:             "Generates GoReleaser's command line docs",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Worker")
			return nil
		},
	}
	root.cmd = cmd
	return root
}
