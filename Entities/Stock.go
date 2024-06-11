package Entities

import "time"

type Stock struct {
	StockID         string
	Quantity        int64
	StockUpdateTime time.Time
}
