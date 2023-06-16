package application

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/uber-go/tally/v4"
	applicationv1 "go.datalift.io/datalift/server/api/application/v1"
	"go.datalift.io/datalift/server/service"
	pgservice "go.datalift.io/datalift/server/service/db/postgres"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/anypb"
)

const Name = "datalift.service.application"

type svc struct {
	logger           *zap.Logger
	scope            tally.Scope
	db               *sql.DB
	advisoryLockConn *sql.Conn
}

type Service interface {
	CreateApplication(ctx context.Context, app *applicationv1.Application) error
	GetApplication(ctx context.Context, id string) (*applicationv1.Application, error)
}

func New(_ *anypb.Any, logger *zap.Logger, scope tally.Scope) (service.Service, error) {
	p, ok := service.Registry[pgservice.Name]
	if !ok {
		return nil, fmt.Errorf("could not find the %v database service", pgservice.Name)
	}

	dbClient, ok := p.(pgservice.Client)
	if !ok {
		return nil, errors.New("database does not implement the required interface")
	}

	return &svc{logger: logger, scope: scope, db: dbClient.DB()}, nil
}

func (s *svc) CreateApplication(ctx context.Context, app *applicationv1.Application) error {
	if app == nil {
		//return nil, errors.New("cannot write empty application to table")
		return errors.New("cannot write empty application to table")
	}

	const writeEventStatement = `INSERT INTO application (name, created_by) VALUES ($1, $2) RETURNING id`
	//err := s.db.QueryRowContext(ctx, writeEventStatement, app.Name, "mberwanger@datalift.com").Scan(&id)
	_, err := s.db.QueryContext(ctx, writeEventStatement, app.Name, `mberwanger@datalift.io`)
	if err != nil {
		//return -1, err
		return err
	}
	//insert into application (name, created_at, created_by) values ('Account Service', now(), 'mberwanger@datalift.io');

	return nil
}

func (s *svc) GetApplication(ctx context.Context, id string) (*applicationv1.Application, error) {
	const stmt = `
		SELECT id, name, created_at, created_by FROM application
		WHERE id = $1
	`

	applications, err := s.query(ctx, stmt, id)
	if err != nil {
		return nil, err
	}
	if len(applications) == 0 {
		return nil, fmt.Errorf("cannot find application by id: %d", id)
	}
	return applications[0], nil
}
