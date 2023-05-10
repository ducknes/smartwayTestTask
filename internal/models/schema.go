package models

type Schema struct {
	Id        uint64   `json:"id"`
	Name      string   `json:"name"`
	Providers []string `json:"providers"`
}
