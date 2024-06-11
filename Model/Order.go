package Model

import "time"

type Order struct {
	OTranID     string    `gorm:"primary_key"`
	OCreated    time.Time `gorm:"DEFAULT:current_timestamp"`
	OStatus     string    `gorm:"NOT NULL"`
	OTotalPrice float64   `gorm:"NOT NULL"`
	ODelivery   bool      `gorm:"NOT NULL"`
}
