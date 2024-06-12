package Entities

import "time"

type Stock struct {
	SID       string
	SQuantity int64
	SUpdated  time.Time
}

func (stock *Stock) Tablename() string {
	return "stocks"
}
