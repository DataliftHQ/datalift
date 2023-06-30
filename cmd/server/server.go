package server

import (
	"github.com/spf13/cobra"

	"go.datalift.io/datalift/cmd/assets"
	"go.datalift.io/datalift/internal/gateway"
)

type ServerCmd struct {
	Cmd  *cobra.Command
	Opts serverOpts
}

type serverOpts struct {
}

func NewServerCmd() *ServerCmd {
	root := &ServerCmd{}
	cmd := &cobra.Command{
		Use:               "server",
		Aliases:           []string{"s"},
		Short:             "Generates GoReleaser's command line docs",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			flags := gateway.ParseFlags()

			components := gateway.CoreComponentFactory
			gateway.Run(flags, components, assets.VirtualFS)

			return nil
		},
	}
	root.Cmd = cmd
	return root
}
