CREATE TABLE app_user (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL
);

CREATE TABLE task_status (
    status VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE task (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(100) NOT NULL REFERENCES task_status(status),
    complete_deadline TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES app_user(id) ON DELETE CASCADE
);

ALTER TABLE app_user ADD CONSTRAINT unique_email UNIQUE (email);