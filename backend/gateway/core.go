package gateway

import (
	"github.com/DataliftHQ/datalift/backend/middleware"
	"github.com/DataliftHQ/datalift/backend/middleware/stats"
	"github.com/DataliftHQ/datalift/backend/middleware/validate"
	"github.com/DataliftHQ/datalift/backend/module"
	assetsmod "github.com/DataliftHQ/datalift/backend/module/assets"
	healthcheckmod "github.com/DataliftHQ/datalift/backend/module/healthcheck"
	"github.com/DataliftHQ/datalift/backend/service"
	awsservice "github.com/DataliftHQ/datalift/backend/service/aws"
	pgservice "github.com/DataliftHQ/datalift/backend/service/db/postgres"
)

var Middleware = middleware.Factory{
	stats.Name:    stats.New,
	validate.Name: validate.New,
}

var Modules = module.Factory{
	assetsmod.Name:      assetsmod.New,
	healthcheckmod.Name: healthcheckmod.New,
}

var Services = service.Factory{
	awsservice.Name: awsservice.New,
	pgservice.Name:  pgservice.New,
}

var CoreComponentFactory = &ComponentFactory{
	Services:   Services,
	Middleware: Middleware,
	Modules:    Modules,
}
