package main

import (
	"github.com/DataliftHQ/datalift/backend/cmd/assets"
	"github.com/DataliftHQ/datalift/backend/gateway"
	"github.com/DataliftHQ/datalift/backend/service"
)

var MockServiceFactory = service.Factory{}

func main() {
	cf := gateway.CoreComponentFactory

	// Replace core services with any available mocks.
	cf.Services = MockServiceFactory

	gateway.Run(gateway.ParseFlags(), cf, assets.VirtualFS)
}
