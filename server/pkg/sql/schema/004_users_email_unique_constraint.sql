-- +goose Up
ALTER TABLE app_user ADD CONSTRAINT unique_email UNIQUE (email);

-- +goose Down
ALTER TABLE app_user DROP CONSTRAINT unique_email;