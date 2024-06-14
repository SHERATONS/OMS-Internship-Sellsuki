package Entities

import "time"

type Address struct {
	City     string
	Country  string
	APrice   float64
	AUpdated time.Time
}

func (address *Address) Tablename() string {
	return "addresses"
}
