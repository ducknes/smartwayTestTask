-- +goose Up
create table airline_provider(
	id serial primary key,
	airline_id int not null,
	provider_id int not null,
	
	foreign key (airline_id) references public.airline(id) on delete cascade,
	foreign key (provider_id) references public.provider(id) on delete cascade
);

-- +goose Down
DROP TABLE IF EXISTS public.airline_provider CASCADE;