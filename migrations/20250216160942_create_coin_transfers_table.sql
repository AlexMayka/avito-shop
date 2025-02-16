-- +goose Up
-- +goose StatementBegin

CREATE TABLE coin_transfers (
    id              SERIAL PRIMARY KEY,
    from_user_id    INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    to_user_id      INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    amount          BIGINT NOT NULL CHECK (amount > 0),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
) WITH (FILLFACTOR = 90);

CREATE INDEX idx_coin_transfers_from_user_id ON coin_transfers(from_user_id);
CREATE INDEX idx_coin_transfers_to_user_id   ON coin_transfers(to_user_id);
CREATE INDEX idx_coin_transfers_created_at   ON coin_transfers(created_at DESC);

CLUSTER coin_transfers USING idx_coin_transfers_to_user_id;
ANALYZE coin_transfers;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS coin_transfers CASCADE;
DROP INDEX IF EXISTS idx_coin_transfers_from_user_id;
DROP INDEX IF EXISTS idx_coin_transfers_to_user_id;
DROP INDEX IF EXISTS idx_coin_transfers_created_at;

-- +goose StatementEnd