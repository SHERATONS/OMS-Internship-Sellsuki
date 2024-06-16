package Entities

import "time"

type Order struct {
	OID          int
	OTranID      string
	OPrice       float64
	ODestination string
	OStatus      string
	OPaid        bool
	OCreated     time.Time
}

func (orders *Order) Tablename() string {
	return "orders"
}
