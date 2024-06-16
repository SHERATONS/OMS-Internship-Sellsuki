package Entities

import "time"

type Order struct {
	OCId         int64
	OCTranID     string
	OCPrice      float64
	ODestination string
	OCStatus     string
	OPaid        bool
	OCreated     time.Time
}

func (orders *Order) Tablename() string {
	return "orders"
}
