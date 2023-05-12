CREATE DATABASE smartwayTestTask;

DROP TABLE IF EXISTS public.airline CASCADE;
DROP TABLE IF EXISTS public.provider CASCADE;
DROP TABLE IF EXISTS public."schema" CASCADE;
DROP TABLE IF EXISTS public.account CASCADE;
DROP TABLE IF EXISTS public.schema_provider CASCADE;
DROP TABLE IF EXISTS public.airline_provider CASCADE;

CREATE TABLE airline(
    id serial primary key,
    iata CHAR(2) NOT null,
    name TEXT NOT NULL
);

CREATE TABLE provider(
    id serial primary key,
    provider_id CHAR(2) NOT NULL,
    name TEXT NOT NULL
);

CREATE TABLE schema(
    id serial PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE account(
    id serial NOT NULL PRIMARY KEY,
    schema_id int not null,
    name TEXT NOT NULL,
    
    foreign key (schema_id) references public."schema"(id)
);

CREATE TABLE schema_provider(
    id serial primary key,
	provider_id int not null, 
	schema_id int not null, 

    foreign key (provider_id) references public.provider(id) on delete cascade,
   	foreign key (schema_id) references public."schema"(id) on delete cascade
);

create table airline_provider(
	id serial primary key,
	airline_id int not null,
	provider_id int not null,
	
	foreign key (airline_id) references public.airline(id) on delete cascade,
	foreign key (provider_id) references public.provider(id) on delete cascade
);