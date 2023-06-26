package application

import (
	"context"
	"errors"
	"github.com/uber-go/tally/v4"
	applicationv1 "go.datalift.io/datalift/api/application/v1"
	"go.datalift.io/datalift/server/module"
	"go.datalift.io/datalift/server/service"
	"go.datalift.io/datalift/server/service/application"
	"go.datalift.io/datalift/server/service/authn"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

const Name = "datalift.module.application"

func New(cfg *anypb.Any, logger *zap.Logger, scope tally.Scope) (module.Module, error) {
	svc, ok := service.Registry["datalift.service.application"]
	if !ok {
		return nil, errors.New("could not find service")
	}

	appService, ok := svc.(application.Service)
	if !ok {
		return nil, errors.New("service was not the correct type")
	}

	m := &mod{
		logger:     logger,
		scope:      scope,
		appService: appService,
	}
	return m, nil
}

type mod struct {
	logger     *zap.Logger
	scope      tally.Scope
	appService application.Service
}

func (m *mod) Register(r module.Registrar) error {
	applicationv1.RegisterApplicationAPIServer(r.GRPCServer(), m)
	return r.RegisterJSONGateway(applicationv1.RegisterApplicationAPIHandler)
}

func (m *mod) CreateApplication(ctx context.Context, req *applicationv1.CreateApplicationRequest) (*applicationv1.CreateApplicationResponse, error) {
	m.logger.Info("CreateApplication")
	err := m.appService.CreateApplication(ctx, req.Application)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	return &applicationv1.CreateApplicationResponse{}, nil
}

func (m *mod) DeleteApplication(ctx context.Context, req *applicationv1.DeleteApplicationRequest) (*applicationv1.DeleteApplicationResponse, error) {
	m.logger.Info("DeleteApplication")
	return nil, nil
}

func (m *mod) GetApplication(ctx context.Context, req *applicationv1.GetApplicationRequest) (*applicationv1.GetApplicationResponse, error) {
	app, err := m.appService.GetApplication(ctx, req.Id)
	if err != nil {

		switch {
		//case errors.Is(err, ErrDivideByZero):
		//fmt.Println("divide by zero error")
		default:
			return nil, status.Errorf(codes.NotFound, "application was not found")
		}
	}

	return &applicationv1.GetApplicationResponse{Application: app}, nil
}

func (m *mod) ListApplications(ctx context.Context, req *applicationv1.ListApplicationsRequest) (*applicationv1.ListApplicationsResponse, error) {
	subject := "Anonymous User" // Used if auth is disabled or it's the actual anonymous user.
	if claims, err := authn.ClaimsFromContext(ctx); err == nil && claims.Subject != authn.AnonymousSubject {
		subject = claims.Subject
	}

	m.logger.Info("ListApplications - " + subject)
	return nil, nil
}

func (m *mod) UpdateApplication(ctx context.Context, req *applicationv1.UpdateApplicationRequest) (*applicationv1.UpdateApplicationResponse, error) {
	m.logger.Info("UpdateApplication")
	return nil, nil
}
