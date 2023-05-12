-- +goose Up
CREATE TABLE schema_provider(
    id serial primary key,
	provider_id int not null, 
	schema_id int not null, 

    foreign key (provider_id) references public.provider(id) on delete cascade,
   	foreign key (schema_id) references public."schema"(id) on delete cascade
);

-- +goose Down
DROP TABLE IF EXISTS public.schema_provider CASCADE;