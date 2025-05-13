-- +goose Up
CREATE TABLE task_status (
    status VARCHAR(100) NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE task_status;