-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
    id          SERIAL PRIMARY KEY,
    username    VARCHAR(255) NOT NULL UNIQUE,
    password    VARCHAR(255) NOT NULL,
    balance     BIGINT       NOT NULL CHECK (balance >= 0),
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
) WITH (FILLFACTOR = 90);

CREATE INDEX idx_users_username ON users(username);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS users CASCADE;
DROP INDEX IF EXISTS idx_users_username;

-- +goose StatementEnd