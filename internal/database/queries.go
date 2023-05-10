package database

var InsertAirline = `INSERT INTO airline (iata, name) VALUES ($1, $2)`

var InsertProvider = `INSERT INTO provider (provider_id, name) VALUES ($1, $2)`

var InsertSchema = `INSERT INTO "schema" (name) VALUES ($1)`

var InsertAccount = `INSERT INTO account (schema_id, name) VALUES ($1, $2)`
