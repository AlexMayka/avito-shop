-- +goose Up
-- +goose StatementBegin

CREATE TABLE merch (
    id     SERIAL       PRIMARY KEY,
    name   VARCHAR(255) NOT NULL UNIQUE,
    price  BIGINT       NOT NULL CHECK (price > 0)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS merch CASCADE;

-- +goose StatementEnd