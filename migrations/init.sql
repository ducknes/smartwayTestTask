CREATE DATABASE smartwayTestTask;

CREATE TABLE airline(
    iata CHAR(2) NOT NULL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE provider(
    id CHAR(2) NOT NULL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE schema(
    id serial NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    providers char(2)[]
);

CREATE TABLE account(
    id serial NOT NULL PRIMARY KEY,
    schemaid INT NOT NULL,
);
