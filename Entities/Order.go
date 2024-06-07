package Entities

import "time"

type Order struct {
	OrderID         int64
	OrderDate       time.Time
	OrderStatus     int
	OrderTotalPrice float64
	OrderDelivery   bool
}
