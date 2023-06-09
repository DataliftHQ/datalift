package validate

import (
	"github.com/golang/protobuf/ptypes/any"
	validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/DataliftHQ/datalift/backend/middleware"
)

const Name = "datalift.middleware.validate"

func New(cfg *any.Any, logger *zap.Logger, scope tally.Scope) (middleware.Middleware, error) {
	return &mid{}, nil
}

type mid struct{}

func (m *mid) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return validator.UnaryServerInterceptor()
}
