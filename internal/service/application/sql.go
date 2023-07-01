package application

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	applicationv1 "go.datalift.io/datalift/api/application/v1"
)

func (s *svc) query(ctx context.Context, query string, args ...interface{}) ([]*applicationv1.Application, error) {
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		s.logger.Error("error querying db", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var applications []*applicationv1.Application
	for rows.Next() {
		row := &applicationv1.Application{}

		var createAt time.Time
		if err := rows.Scan(&row.Id, &row.Name, &createAt, &row.CreatedBy); err != nil {
			s.logger.Error("error scanning db results", zap.Error(err))
			return nil, err
		}

		createAtProto := timestamppb.New(createAt)
		if err := createAtProto.CheckValid(); err != nil {
			s.logger.Error("error in parsing db result's timestamp", zap.Error(err))
			return nil, err
		}
		row.CreatedAt = createAtProto

		applications = append(applications, row)
	}

	return applications, nil
}
