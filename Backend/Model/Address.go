package Model

import "time"

type Address struct {
	City     string  `gorm:"primary_key"`
	Country  string  `gorm:"NOT NULL"`
	APrice   float64 `gorm:"NOT NULL"`
	AUpdated time.Time
}
