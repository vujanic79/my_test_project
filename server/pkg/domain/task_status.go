package domain

import (
	"context"
	"net/http"
)

type TaskStatus struct {
	Status string `json:"status"`
}

type CreateTaskStatusParams struct {
	Status string `json:"status"`
}

type TaskStatusService interface {
	CreateTaskStatus(ctx context.Context, status string) (ts TaskStatus, err error)
	GetTaskStatuses(ctx context.Context) (tss []TaskStatus, err error)
	GetTaskStatusByStatus(ctx context.Context, status string) (ts TaskStatus, err error)
}

type TaskStatusRepository interface {
	CreateTaskStatus(ctx context.Context, status string) (ts TaskStatus, err error)
	GetTaskStatuses(ctx context.Context) (tss []TaskStatus, err error)
	GetTaskStatusByStatus(ctx context.Context, status string) (ts TaskStatus, err error)
}

type TaskStatusController interface {
	CreateTaskStatus(w http.ResponseWriter, r *http.Request)
	GetTaskStatuses(w http.ResponseWriter, r *http.Request)
	GetTaskStatusByStatus(w http.ResponseWriter, r *http.Request)
}
