package Entities

import "time"

type Stock struct {
	SID       string
	SQuantity int8
	SUpdated  time.Time
}

func (stock *Stock) Tablename() string {
	return "stocks"
}
