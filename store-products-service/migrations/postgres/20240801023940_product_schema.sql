-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    title VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    price VARCHAR NOT NULL,
    category VARCHAR NOT NULL,
    quantity VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
