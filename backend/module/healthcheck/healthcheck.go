package healthcheck

import (
	"context"

	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/anypb"

	healthcheckv1 "github.com/DataliftHQ/datalift/backend/api/healthcheck/v1"
	"github.com/DataliftHQ/datalift/backend/module"
)

const Name = "datalift.module.healthcheck"

func New(*anypb.Any, *zap.Logger, tally.Scope) (module.Module, error) {
	m := &mod{
		api: newAPI(),
	}
	return m, nil
}

type mod struct {
	api healthcheckv1.HealthcheckAPIServer
}

func (m *mod) Register(r module.Registrar) error {
	healthcheckv1.RegisterHealthcheckAPIServer(r.GRPCServer(), m.api)
	return r.RegisterJSONGateway(healthcheckv1.RegisterHealthcheckAPIHandler)
}

func newAPI() healthcheckv1.HealthcheckAPIServer {
	return &healthcheckAPI{}
}

type healthcheckAPI struct{}

func (a *healthcheckAPI) Healthcheck(context.Context, *healthcheckv1.HealthcheckRequest) (*healthcheckv1.HealthcheckResponse, error) {
	return &healthcheckv1.HealthcheckResponse{}, nil
}
