package gateway

import (
	"github.com/DataliftHQ/datalift/backend/middleware"
	"github.com/DataliftHQ/datalift/backend/middleware/authn"
	"github.com/DataliftHQ/datalift/backend/middleware/stats"
	"github.com/DataliftHQ/datalift/backend/middleware/validate"
	"github.com/DataliftHQ/datalift/backend/module"
	appmod "github.com/DataliftHQ/datalift/backend/module/application"
	assetsmod "github.com/DataliftHQ/datalift/backend/module/assets"
	authnmod "github.com/DataliftHQ/datalift/backend/module/authn"
	healthcheckmod "github.com/DataliftHQ/datalift/backend/module/healthcheck"
	"github.com/DataliftHQ/datalift/backend/service"
	appservice "github.com/DataliftHQ/datalift/backend/service/application"
	authnservice "github.com/DataliftHQ/datalift/backend/service/authn"
	awsservice "github.com/DataliftHQ/datalift/backend/service/aws"
	pgservice "github.com/DataliftHQ/datalift/backend/service/db/postgres"
	workflowservice "github.com/DataliftHQ/datalift/backend/service/workflow"
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
