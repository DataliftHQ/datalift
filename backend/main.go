package main

import (
	"github.com/DataliftHQ/datalift/backend/cmd/assets"
	"github.com/DataliftHQ/datalift/backend/gateway"
)

func main() {
	flags := gateway.ParseFlags()
	components := gateway.CoreComponentFactory

	gateway.Run(flags, components, assets.VirtualFS)
}
