package Entities

type Address struct {
	AID     string
	City    string
	Country string
}

func (address *Address) Tablename() string {
	return "addresses"
}
