package main

import (
	"go.datalift.io/datalift/server/cmd/assets"
	"go.datalift.io/datalift/server/gateway"
)

//nolint:all
var (
	version = ""
)

func main() {
	flags := gateway.ParseFlags()
	components := gateway.CoreComponentFactory

	gateway.Run(flags, components, assets.VirtualFS)
}
