package Entities

import "time"

type Stock struct {
	SID       string
	SQuantity float64
	SUpdated  time.Time
}