package gateway

import (
	"go.datalift.io/datalift/internal/middleware"
	"go.datalift.io/datalift/internal/middleware/authn"
	"go.datalift.io/datalift/internal/middleware/stats"
	"go.datalift.io/datalift/internal/middleware/validate"
	"go.datalift.io/datalift/internal/module"
	appmod "go.datalift.io/datalift/internal/module/application"
	assetsmod "go.datalift.io/datalift/internal/module/assets"
	authnmod "go.datalift.io/datalift/internal/module/authn"
	healthcheckmod "go.datalift.io/datalift/internal/module/healthcheck"
	"go.datalift.io/datalift/internal/service"
	appservice "go.datalift.io/datalift/internal/service/application"
	authnservice "go.datalift.io/datalift/internal/service/authn"
	awsservice "go.datalift.io/datalift/internal/service/aws"
	pgservice "go.datalift.io/datalift/internal/service/db/postgres"
	workflowservice "go.datalift.io/datalift/internal/service/workflow"
)

var Middleware = middleware.Factory{
	stats.Name:    stats.New,
	authn.Name:    authn.New,
	validate.Name: validate.New,
}

// TODO: Modules is a terrible name
var Modules = module.Factory{
	assetsmod.Name:      assetsmod.New,
	authnmod.Name:       authnmod.New,
	healthcheckmod.Name: healthcheckmod.New,
	appmod.Name:         appmod.New,
}

var Services = service.Factory{
	authnservice.Name:        authnservice.New,
	authnservice.StorageName: authnservice.NewStorage,
	awsservice.Name:          awsservice.New,
	pgservice.Name:           pgservice.New,
	appservice.Name:          appservice.New,
	workflowservice.Name:     workflowservice.New,
}

var CoreComponentFactory = &ComponentFactory{
	Services:   Services,
	Middleware: Middleware,
	Modules:    Modules,
}
