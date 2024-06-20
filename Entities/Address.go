package Entities

import "time"

type Address struct {
	City     string
	Country  string
	APrice   float64
	AUpdated time.Time
}