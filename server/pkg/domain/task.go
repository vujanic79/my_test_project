package domain

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

type Task struct {
	ID               uuid.UUID `json:"id"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Status           string    `json:"status"`
	CompleteDeadline time.Time `json:"completeDeadline"`
	UserID           uuid.UUID `json:"userId"`
}

type CreateTaskParams struct {
	ID               uuid.UUID `json:"id"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Status           string    `json:"status"`
	CompleteDeadline string    `json:"completeDeadline"`
	UserEmail        string    `json:"userEmail"`
}

func (ctp CreateTaskParams) MarshalZerologObject(event *zerolog.Event) {
	event.
		Str("title", ctp.Title).
		Str("description", ctp.Description).
		Str("status", ctp.Status).
		Str("completeDeadline", ctp.CompleteDeadline).
		Str("userEmail", ctp.UserEmail)
}

type UpdateTaskParams struct {
	ID               uuid.UUID `json:"id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Status           string    `json:"status"`
	CompleteDeadline string    `json:"completeDeadline"`
}

func (utp UpdateTaskParams) MarshalZerologObject(e *zerolog.Event) {
	e.
		Interface("id", utp.ID).
		Str("title", utp.Title).
		Str("description", utp.Description).
		Str("status", utp.Status).
		Str("completeDeadline", utp.CompleteDeadline)
}

type GetTasksByUserIdParams struct {
	UserID uuid.UUID `json:"userId"`
}

type TaskService interface {
	CreateTask(ctx context.Context, userId uuid.UUID, params CreateTaskParams) (t Task, err error)
	DeleteTask(ctx context.Context, id uuid.UUID) (err error)
	UpdateTask(ctx context.Context, params UpdateTaskParams) (t Task, err error)
	GetTasksByUserId(ctx context.Context, id uuid.UUID) (ts []Task, err error)
}

type TaskRepository interface {
	CreateTask(ctx context.Context, userId uuid.UUID, params CreateTaskParams) (t Task, err error)
	DeleteTask(ctx context.Context, id uuid.UUID) (err error)
	UpdateTask(ctx context.Context, params UpdateTaskParams) (t Task, err error)
	GetTasksByUserId(ctx context.Context, id uuid.UUID) (ts []Task, err error)
}

type TaskController interface {
	CreateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	GetTasksByUserId(w http.ResponseWriter, r *http.Request)
}
