package service

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
)

type TaskStatusService struct {
	Tsr domain.TaskStatusRepository
}

var _ domain.TaskStatusService = (*TaskStatusService)(nil)

func NewTaskStatusService(tsr domain.TaskStatusRepository) (tss *TaskStatusService) {
	return &TaskStatusService{Tsr: tsr}
}

func (tss *TaskStatusService) CreateTaskStatus(ctx context.Context, status string) (ts domain.TaskStatus, err error) {
	l := logger.FromContext(ctx)

	l = l.With().
		Dict("service.params", zerolog.Dict().
			Str("func", "CreateTaskStatus").
			Dict("params", zerolog.Dict().
				Str("status", status))).
		Logger()
	ctx = logger.WithLogger(ctx, l)

	ts, err = tss.Tsr.CreateTaskStatus(ctx, status)
	return ts, err
}

func (tss *TaskStatusService) GetTaskStatuses(ctx context.Context) (taskStatuses []domain.TaskStatus, err error) {
	l := logger.FromContext(ctx)

	l = l.With().
		Dict("service.params", zerolog.Dict().
			Str("func", "GetTaskStatuses").
			Dict("params", zerolog.Dict())).
		Logger()
	ctx = logger.WithLogger(ctx, l)

	taskStatuses, err = tss.Tsr.GetTaskStatuses(ctx)
	return taskStatuses, err
}

func (tss *TaskStatusService) GetTaskStatusByStatus(ctx context.Context, status string) (ts domain.TaskStatus, err error) {
	l := logger.FromContext(ctx)

	l = l.With().
		Dict("service.params", zerolog.Dict().
			Str("func", "GetTaskStatusByStatus").
			Dict("params", zerolog.Dict().
				Str("status", status))).
		Logger()
	ctx = logger.WithLogger(ctx, l)

	ts, err = tss.Tsr.GetTaskStatusByStatus(ctx, status)
	return ts, err
}
