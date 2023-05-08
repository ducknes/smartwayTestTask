package models

type provider struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewProvider(id string, name string) *provider {
	return &provider{
		Id:   id,
		Name: name,
	}
}
