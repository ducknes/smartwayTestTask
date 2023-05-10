package models

type Account struct {
	Id       uint64 `json:"id"`
	SchemaId uint64 `json:"schema_id"`
	Name     string `json:"name"`
}
