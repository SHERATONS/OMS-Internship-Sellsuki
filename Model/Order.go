package Model

import "time"

type Order struct {
	OCId         int64     `gorm:"primary_key"`
	OCTranID     string    `gorm:"not null"`
	OCPrice      float64   `gorm:"not null"`
	ODestination string    `gorm:"not null"`
	OCStatus     string    `gorm:"not null"`
	OPaid        bool      `gorm:"not null"`
	OCreated     time.Time `gorm:"not null"`
}
