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

-- insert into public.airline (iata, name) values ('SU','Аэрофлот');
-- insert into public.airline (iata, name) values ('S7','S7');
-- insert into public.airline (iata, name) values ('KV','КрасАвиа');
-- insert into public.airline (iata, name) values ('U6','Уральские авиалинии');
-- insert into public.airline (iata, name) values ('UT','ЮТэйр');
-- insert into public.airline (iata, name) values ('FZ','Flydubai');
-- insert into public.airline (iata, name) values ('JB','JetBlue');
-- insert into public.airline (iata, name) values ('SJ','SuperJet');
-- insert into public.airline (iata, name) values ('WZ','Wizz Air');
-- insert into public.airline (iata, name) values ('N4','Nordwind Airlines');
-- insert into public.airline (iata, name) values ('5N','SmartAvia');

-- insert into public.provider (id, name) values ('AA', 'AmericanAir');
-- insert into public.provider (id, name) values ('IF', 'InternationFlights');
-- insert into public.provider (id, name) values ('RS', 'RedStar');

-- insert into public."schema" (name) values ('Основная');
-- insert into public."schema" (name) values ('Тестовая');

-- insert into public."schema" (name, providers) values ('Основная', (array(select id from public.provider)));
-- insert into public."schema" (name, providers) values ('Тестовая', (array(select id from public.provider where id!='AA')));


-- update public."schema" set providers=array(select id from public.provider where id!='AA') where name='тестовая';
-- update public."schema" set providers=array(select id from public.provider) where name='основная';

-- insert into public.account (schemaid, name) values ((select id from public."schema" where name='Тестовая'), 'Демо');
-- insert into public.account (schemaid, name) values ((select id from public."schema" where name='Тестовая'), 'Разработка');
-- insert into public.account (schemaid, name) values ((select id from public."schema" where name='Основная'), 'Продажи');

-- insert into public.schema_prvider (providerid, schemaid) values ('AA','1');
-- insert into public.schema_prvider (providerid, schemaid) values ('IF','1');
-- insert into public.schema_prvider (providerid, schemaid) values ('RS','1');
-- insert into public.schema_prvider (providerid, schemaid) values ('IF','2');
-- insert into public.schema_prvider (providerid, schemaid) values ('RS','2');

-- update public."schema" p set providers=array(select ps.providerid from public.schema_prvider ps where ps.schemaid=p.id);
-- update public."schema" p set providers=array(select ps.providerid from public.schema_prvider ps where ps.schemaid=p.id);

-- SELECT provider.id, schema.name
-- FROM provider
-- JOIN public.schema_prvider ON provider.id = schema_prvider.providerid
-- JOIN public."schema" ON schema_prvider.schemaid = schema.id;