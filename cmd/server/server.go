package server

import (
	"fmt"
	"github.com/spf13/cobra"
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

			//package main
			//
			//import (
			//	"go.datalift.io/datalift/internal/gateway"
			//	"go.datalift.io/datalift/server/cmd/assets"
			//)
			//
			////nolint:all
			//var (
			//	version = ""
			//)
			//
			//func main() {
			//	flags := gateway.ParseFlags()
			//	components := gateway.CoreComponentFactory
			//
			//	gateway.Run(flags, components, assets.VirtualFS)
			//}
			fmt.Println("server")
			return nil
		},
	}
	root.Cmd = cmd
	return root
}
