package models

import "fmt"

type Airline struct {
	Id   uint64 `json:"id"`
	Iata string `json:"iata"`
	Name string `json:"name"`
}

func (a *Airline) String() string {
	return fmt.Sprintf("id : %d, iata : %s, name : %s", a.Id, a.Iata, a.Name)
}
