package models

type providerID string

type schema struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	Providers []providerID `json:"providers"`
}

func newSchema(id uint64, name string, providers []providerID) *schema {
	return &schema{
		Id:        id,
		Name:      name,
		Providers: providers,
	}
}
