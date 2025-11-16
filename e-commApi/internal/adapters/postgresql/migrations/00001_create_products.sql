-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query'; -- up -> what we are gonna change
CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price_in_centers INTEGER NOT NULL CHECK (price_in_centers>=0),
    quantity INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query'; -- down->how we revert the change
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
