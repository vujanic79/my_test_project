package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
)

type TaskService struct {
	Tr domain.TaskRepository
}

var _ domain.TaskService = (*TaskService)(nil)

func NewTaskService(tr domain.TaskRepository) (ts TaskService) {
	return TaskService{Tr: tr}
}

func (ts TaskService) CreateTask(ctx context.Context, userId uuid.UUID, params domain.CreateTaskParams) (t domain.Task, err error) {
	l := logger.FromContext(ctx)

	l = l.With().
		Dict("service.params", zerolog.Dict().
			Str("func", "CreateTask").
			Dict("params", zerolog.Dict().
				Interface("userId", userId).
				Object("params", params))).
		Logger()
	ctx = logger.WithLogger(ctx, l)

	return ts.Tr.CreateTask(ctx, userId, params)
}

func (ts TaskService) DeleteTask(ctx context.Context, id uuid.UUID) (err error) {
	l := logger.FromContext(ctx)

	l = l.With().
		Dict("service.params", zerolog.Dict().
			Str("func", "DeleteTask").
			Dict("params", zerolog.Dict().
				Interface("id", id))).
		Logger()
	ctx = logger.WithLogger(ctx, l)

	return ts.Tr.DeleteTask(ctx, id)
}

func (ts TaskService) UpdateTask(ctx context.Context, params domain.UpdateTaskParams) (t domain.Task, err error) {
	l := logger.FromContext(ctx)

	l = l.With().
		Dict("service.params", zerolog.Dict().
			Str("func", "UpdateTask").
			Object("params", params)).
		Logger()
	ctx = logger.WithLogger(ctx, l)

	return ts.Tr.UpdateTask(ctx, params)
}

func (ts TaskService) GetTasksByUserId(ctx context.Context, id uuid.UUID) (tasks []domain.Task, err error) {
	l := logger.FromContext(ctx)

	l = l.With().
		Dict("service.params", zerolog.Dict().
			Str("func", "GetTasksByUserId").
			Dict("params", zerolog.Dict().
				Interface("userId", id))).
		Logger()
	ctx = logger.WithLogger(ctx, l)

	return ts.Tr.GetTasksByUserId(ctx, id)
}
