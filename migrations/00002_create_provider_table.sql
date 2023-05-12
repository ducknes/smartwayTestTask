-- +goose Up
CREATE TABLE provider(
    id serial primary key,
    provider_id CHAR(2) NOT NULL,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS public.provider CASCADE;