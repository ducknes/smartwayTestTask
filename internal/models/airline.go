package models

type Airline struct {
	Iata string `json:"iata"`
	Name string `json:"name"`
}

func NewAirline(iata string, name string) *Airline {
	return &Airline{
		Iata: iata,
		Name: name,
	}
}
