package Entities

type Address struct {
	AddressID int64  `json:"addressId"`
	City      string `json:"city"`
	Country   string `json:"country"`
}
