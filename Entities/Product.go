package Entities

import "time"

type Product struct {
	PID      string
	PName    string
	PPrice   float64
	PDesc    string
	PCreated time.Time
	PUpdated time.Time
}