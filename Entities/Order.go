package Entities

import "time"

type Order struct {
	OTranID     string
	OCreated    time.Time
	OStatus     string
	OTotalPrice float64
	ODelivery   bool
}

func (order *Order) Tablename() string {
	return "orders"
}
