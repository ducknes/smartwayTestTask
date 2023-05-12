-- +goose Up
CREATE TABLE account(
    id serial NOT NULL PRIMARY KEY,
    schema_id int not null,
    name TEXT NOT NULL,
    
    foreign key (schema_id) references public."schema"(id)
);

-- +goose Down
DROP TABLE IF EXISTS public.account CASCADE;