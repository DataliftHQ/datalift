package application

import (
	"context"
	applicationv1 "go.datalift.io/datalift/server/api/application/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

const createOrUpdateSubmissionQuery = `
INSERT INTO feedback (client_id, submitted_at, user_id, score, details, metadata) VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (client_id) DO UPDATE SET
		client_id = EXCLUDED.client_id,
		submitted_at = EXCLUDED.submitted_at,
		user_id = EXCLUDED.user_id,
		score = EXCLUDED.score,
		details = EXCLUDED.details,
		metadata = EXCLUDED.metadata
`

//
//func (s *storage) createOrUpdateSubmission(ctx context.Context, submission *submission) error {
//	feedbackJSON, err := protojson.Marshal(submission.feedback)
//	if err != nil {
//		return err
//	}
//	metadataJSON, err := protojson.Marshal(submission.metadata)
//	if err != nil {
//		return err
//	}
//	_, err = s.db.ExecContext(ctx, createOrUpdateSubmissionQuery, submission.id, submission.submittedAt, submission.userId, submission.score, feedbackJSON, metadataJSON)
//	return err
//}

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
