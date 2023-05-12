-- +goose Up
CREATE TABLE schema(
    id serial PRIMARY KEY,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS public."schema" CASCADE;
