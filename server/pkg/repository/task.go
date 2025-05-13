package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/internal/database"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
	"strings"
	"time"
)

type TaskRepository struct {
	Db *database.Queries
}

var _ domain.TaskRepository = (*TaskRepository)(nil)

func NewTaskRepository(db *database.Queries) (ts *TaskRepository) {
	return &TaskRepository{Db: db}
}

func (tr *TaskRepository) CreateTask(
	ctx context.Context,
	userId uuid.UUID,
	params domain.CreateTaskParams) (t domain.Task, err error) {
	l := logger.FromContext(ctx)
	layout := "2006-01-02T15:04:05.999999Z"
	parsedTime, err := time.Parse(layout, params.CompleteDeadline)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("repository.params", zerolog.Dict().
				Str("func", "CreateTask").
				Dict("params", zerolog.Dict().
					Str("completeDeadline", params.CompleteDeadline))).
			Msg("Parsing completeDeadline error")
		return domain.Task{}, err
	}

	dbT, err := tr.Db.CreateTask(ctx, database.CreateTaskParams{
		ID:               uuid.New(),
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
		Title:            params.Title,
		Description:      params.Description,
		Status:           strings.ToUpper(params.Status),
		CompleteDeadline: parsedTime,
		UserID:           userId,
	})

	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("repository.params", zerolog.Dict().
				Str("func", "CreateTask").
				Dict("params", zerolog.Dict().
					Interface("userId", userId).
					Object("params", params))).
			Msg("Creating task error")
		return domain.Task{}, err
	}

	return MapDbTaskToTask(dbT), err
}

func (tr *TaskRepository) DeleteTask(ctx context.Context, id uuid.UUID) (err error) {
	l := logger.FromContext(ctx)
	err = tr.Db.DeleteTask(ctx, id)

	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("repository.params", zerolog.Dict().
				Str("func", "DeleteTask").
				Dict("params", zerolog.Dict().
					Interface("id", id))).
			Msg("Deleting task error")
	}

	return err
}

func (tr *TaskRepository) UpdateTask(ctx context.Context, params domain.UpdateTaskParams) (t domain.Task, err error) {
	l := logger.FromContext(ctx)
	layout := "2006-01-02T15:04:05.999999Z"
	parsedTime, err := time.Parse(layout, params.CompleteDeadline)

	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("repository.params", zerolog.Dict().
				Str("func", "UpdateTask").
				Dict("params", zerolog.Dict().
					Str("completeDeadline", params.CompleteDeadline))).
			Msg("Parsing completeDeadline error")
		return domain.Task{}, err
	}

	dbT, err := tr.Db.UpdateTask(ctx, database.UpdateTaskParams{
		ID:               params.ID,
		Title:            params.Title,
		Description:      params.Description,
		CompleteDeadline: parsedTime,
		Status:           strings.ToUpper(params.Status),
	})

	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("repository.params", zerolog.Dict().
				Str("func", "UpdateTask").
				Dict("params", zerolog.Dict().
					Object("params", params))).
			Msg("Updating task error")
		return domain.Task{}, err
	}

	return MapDbTaskToTask(dbT), err
}
func (tr *TaskRepository) GetTasksByUserId(ctx context.Context, id uuid.UUID) (ts []domain.Task, err error) {
	l := logger.FromContext(ctx)
	dbTs, err := tr.Db.GetTasksByUserId(ctx, id)

	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("repository.params", zerolog.Dict().
				Str("func", "GetTasksByUserId").
				Dict("params", zerolog.Dict().
					Interface("userId", id))).
			Msg("Getting tasks by userId error")
		return nil, err
	}

	return MapDbTasksToTasks(dbTs), err
}

type GetTasksByUserIdParams struct {
	UserID uuid.UUID `json:"user_id"`
}

func MapDbTaskToTask(dbT database.Task) (t domain.Task) {
	return domain.Task{
		ID:               dbT.ID,
		CreatedAt:        dbT.CreatedAt,
		UpdatedAt:        dbT.UpdatedAt,
		Title:            dbT.Title,
		Description:      dbT.Description,
		Status:           dbT.Status,
		CompleteDeadline: dbT.CompleteDeadline,
		UserID:           dbT.UserID,
	}
}

func MapDbTasksToTasks(dbTs []database.Task) (ts []domain.Task) {
	ts = make([]domain.Task, len(dbTs))
	for i, dbTask := range dbTs {
		ts[i] = MapDbTaskToTask(dbTask)
	}
	return ts
}
