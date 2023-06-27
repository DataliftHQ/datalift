package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type serverCmd struct {
	cmd  *cobra.Command
	opts serverOpts
}

type serverOpts struct {
}

func newServerCmd() *serverCmd {
	root := &serverCmd{}
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
	root.cmd = cmd
	return root
}
