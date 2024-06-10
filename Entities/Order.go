package Entities

import "time"

type Order struct {
	OrderTranID     string
	OrderCreateTime time.Time
	OrderStatus     string
	OrderTotalPrice float64
	OrderDelivery   bool
}
