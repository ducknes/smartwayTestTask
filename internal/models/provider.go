package models

type Provider struct {
	Id          uint   `json:"id"`
	ProdviderId string `json:"provider_id"`
	Name        string `json:"name"`
}
