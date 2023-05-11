package database

var InsertAirline = `INSERT INTO airline (iata, name) VALUES ($1, $2)`

var InsertProvider = `INSERT INTO provider (provider_id, name) VALUES ($1, $2)`

var InsertSchema = `INSERT INTO "schema" (name) VALUES ($1)`

var InsertAccount = `INSERT INTO account (schema_id, name) VALUES ($1, $2)`

var DeleteAirline = `DELETE FROM airline where iata = $1`

var DeleteProvider = `DELETE FROM provider where provider_id = $1`

var DeleteSchema = `DELETE FROM public."schema" WHERE id = $1 AND NOT EXISTS (SELECT 1 FROM public.account WHERE schema_id = $1)`

var DeleteAccount = `DELETE FROM account WHERE id = $1`

