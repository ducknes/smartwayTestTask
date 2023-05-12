package database

var InsertAirline = `INSERT INTO airline (iata, name) VALUES ($1, $2)`

var InsertProvider = `INSERT INTO provider (provider_id, name) VALUES ($1, $2)`

var InsertSchema = `INSERT INTO "schema" (name) VALUES ($1)`

var InsertAccount = `INSERT INTO account (schema_id, name) VALUES ($1, $2)`

var DeleteAirline = `DELETE FROM airline where iata = $1`

var DeleteProvider = `DELETE FROM provider where provider_id = $1`

var DeleteSchema = `DELETE FROM public."schema" WHERE id = $1 AND NOT EXISTS (SELECT 1 FROM public.account WHERE schema_id = $1)`

var DeleteAccount = `DELETE FROM account WHERE id = $1`

var InsertToAirlineProvires = `INSERT INTO airline_provider (airline_id, provider_id) values ($1, $2)`

var DeleteFromAirlineProvider = `DELETE FROM airline_provider WHERE airline_id = $1`

var SelectAirlineId = `SELECT id FROM airline WHERE iata = $1`

var SelectProviderId = `SELECT id FROM provider WHERE provider_id = $1`

var FindSchemaByName = `SELECT id, name FROM public."schema" WHERE name = $1`

var GetSchemaProviders = `
	SELECT s.name AS schemaname, p.provider_id AS providerid
	FROM schema s
	LEFT JOIN schema_provider sp ON sp.schema_id = s.id
	LEFT JOIN provider p ON p.id = sp.provider_id
	WHERE s.name = $1
`

var UpdateSchemaName = `UPDATE public."schema" SET name = $1 WHERE id = $2`

var InsertToSchemaProvires = `INSERT INTO schema_provider (provider_id, schema_id) values ($1, $2)`

var DeleteFromSchemaProvider = `DELETE FROM schema_provider WHERE schema_id = $1`

var UpdateAccountSchema = `UPDATE account SET schema_id = $1 WHERE id = $2`

var GetAccountAirlines = `
	SELECT a.iata AS airline_id, p.provider_id AS provider_id
	FROM airline a
	LEFT JOIN airline_provider ap ON ap.airline_id = a.id
	LEFT JOIN provider p ON p.id = ap.provider_id
	WHERE p.id = ANY(select provider_id from schema_provider where schema_id IN (select schema_id from account where id=$1))
`

var GetProviderAirlines = `
	SELECT a.iata AS airline_id, p.provider_id AS providerid
	FROM airline a
	LEFT JOIN airline_provider ap ON ap.airline_id = a.id
	LEFT JOIN provider p ON p.id = ap.provider_id
	WHERE p.provider_id = $1
`
