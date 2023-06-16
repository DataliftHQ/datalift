package modulemock

import (
	"google.golang.org/grpc"

	mod "go.datalift.io/datalift/server/module"
)

type MockRegistrar struct {
	Server *grpc.Server
}

func (m *MockRegistrar) GRPCServer() *grpc.Server { return m.Server }

func (m *MockRegistrar) RegisterJSONGateway(handlerFunc mod.GatewayRegisterAPIHandlerFunc) error {
	return nil
}
