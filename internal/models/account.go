package models

type schemaId uint64

type account struct {
	Id       uint64   `json:"id"`
	SchemaId schemaId `json:"schemaId,omitempty"`
}

func NewAccount(id uint64, schemaId schemaId) *account {
	return &account{
		Id:       id,
		SchemaId: schemaId,
	}
}
