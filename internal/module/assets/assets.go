package assets

import (
	"context"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"

	assetsv1 "go.datalift.io/datalift/api/assets/v1"
	"go.datalift.io/datalift/internal/module"
)

const Name = "datalift.module.assets"

func New(*any.Any, *zap.Logger, tally.Scope) (module.Module, error) {
	mod := &mod{
		api: newAPI(),
	}
	return mod, nil
}

type mod struct {
	api assetsv1.AssetsAPIServer
}

func (m *mod) Register(r module.Registrar) error {
	assetsv1.RegisterAssetsAPIServer(r.GRPCServer(), m.api)
	return r.RegisterJSONGateway(assetsv1.RegisterAssetsAPIHandler)
}

func newAPI() assetsv1.AssetsAPIServer {
	return &api{}
}

type api struct{}

func (a *api) Fetch(ctx context.Context, request *assetsv1.FetchRequest) (*assetsv1.FetchResponse, error) {
	return &assetsv1.FetchResponse{}, nil
}
