-- name: CreateTaskStatus :one
INSERT INTO task_status (status)
VALUES ($1)
RETURNING *;

-- name: GetTaskStatuses :many
SELECT * FROM task_status;

-- name: GetTaskStatusByStatus :one
SELECT * FROM task_status ts
WHERE ts.status = $1;