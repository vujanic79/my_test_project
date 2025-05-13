package repository

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/internal/database"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
)

type TaskStatusRepository struct {
	Db *database.Queries
}

var _ domain.TaskStatusRepository = (*TaskStatusRepository)(nil)

func NewTaskStatusRepository(db *database.Queries) (tsr *TaskStatusRepository) {
	return &TaskStatusRepository{Db: db}
}

func (tsr *TaskStatusRepository) CreateTaskStatus(ctx context.Context, status string) (ts domain.TaskStatus, err error) {
	l := logger.FromContext(ctx)
	dbStatus, err := tsr.Db.CreateTaskStatus(ctx, status)

	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("repository.params", zerolog.Dict().
				Str("func", "CreateTaskStatus").
				Dict("params", zerolog.Dict().
					Str("status", status))).
			Msg("Creating task status error")
		return domain.TaskStatus{}, err
	}

	return domain.TaskStatus{Status: dbStatus}, err
}

func (tsr *TaskStatusRepository) GetTaskStatuses(ctx context.Context) (tss []domain.TaskStatus, err error) {
	l := logger.FromContext(ctx)
	dbTss, err := tsr.Db.GetTaskStatuses(ctx)

	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("repository.params", zerolog.Dict().
				Str("func", "GetTaskStatuses").
				Dict("params", zerolog.Dict())).
			Msg("Getting task statuses from database error")
		return nil, err
	}

	return MapDbTaskStatusesToTaskStatuses(dbTss), err
}

func (tsr *TaskStatusRepository) GetTaskStatusByStatus(ctx context.Context, status string) (ts domain.TaskStatus, err error) {
	l := logger.FromContext(ctx)
	dbTs, err := tsr.Db.GetTaskStatusByStatus(ctx, status)

	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("repository.params", zerolog.Dict().
				Str("func", "GetTaskStatusByStatus").
				Dict("params", zerolog.Dict().
					Str("status", status))).
			Msg("Getting task status from database error")
		return domain.TaskStatus{}, err
	}

	return MapDbTaskStatusToTaskStatus(dbTs), err
}

func MapDbTaskStatusToTaskStatus(dbTs string) (ts domain.TaskStatus) {
	return domain.TaskStatus{Status: dbTs}
}

func MapDbTaskStatusesToTaskStatuses(dbTss []string) (tss []domain.TaskStatus) {
	tss = make([]domain.TaskStatus, len(dbTss))
	for i, taskStatus := range dbTss {
		tss[i] = MapDbTaskStatusToTaskStatus(taskStatus)
	}
	return tss
}
