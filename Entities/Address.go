package Entities

type Address struct {
	AID     string
	City    string
	Country string
	APrice  float64
}

func (address *Address) Tablename() string {
	return "addresses"
}
