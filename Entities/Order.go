package Entities

import "time"

type Order struct {
	OrderTranID     int64     `json:"orderTranID"`
	OrderCreateTime time.Time `json:"orderCreateTime"`
	OrderStatus     int       `json:"orderStatus"`
	OrderTotalPrice float64   `json:"orderTotalPrice"`
	OrderDelivery   bool      `json:"orderDelivery"`
}
