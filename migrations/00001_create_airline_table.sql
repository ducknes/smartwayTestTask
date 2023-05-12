-- +goose Up
CREATE TABLE airline(
    id serial primary key,
    iata CHAR(2) NOT null,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS public.airline CASCADE;