-- +goose Up
-- +goose StatementBegin

CREATE TABLE purchases (
    id            SERIAL PRIMARY KEY,
    user_id       INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    merch_id      INT NOT NULL REFERENCES merch(id) ON DELETE CASCADE,
    price_bought  BIGINT NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
) WITH (FILLFACTOR = 90);

CREATE INDEX idx_purchases_user_id ON purchases(user_id);
CREATE INDEX idx_purchases_merch_id ON purchases(merch_id);
CREATE INDEX idx_purchases_created_at ON purchases(created_at);

CLUSTER purchases USING idx_purchases_user_id;
ANALYZE purchases;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS purchases CASCADE;
DROP INDEX IF EXISTS idx_purchases_user_id;
DROP INDEX IF EXISTS idx_purchases_merch_id;
DROP INDEX IF EXISTS idx_purchases_created_at;

-- +goose StatementEnd