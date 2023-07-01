package main

import (
	"go.datalift.io/datalift/cmd/assets"
	"go.datalift.io/datalift/internal/gateway"
	"go.datalift.io/datalift/internal/service"
)

var MockServiceFactory = service.Factory{}

func main() {
	cf := gateway.CoreComponentFactory

	// Replace core services with any available mocks.
	cf.Services = MockServiceFactory

	gateway.Run(gateway.ParseFlags(), cf, assets.VirtualFS)
}
