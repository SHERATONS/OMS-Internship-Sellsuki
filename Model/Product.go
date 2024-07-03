package Model

import "time"

type Product struct {
	PID      string    `gorm:"primary_key"`
	PName    string    `gorm:"NOT NULL"`
	PPrice   float64   `gorm:"NOT NULL"`
	PDesc    string    `gorm:"NOT NULL"`
	PCreated time.Time `gorm:"DEFAULT:current_timestamp"`
	PUpdated time.Time `gorm:"DEFAULT:current_timestamp"`
}
