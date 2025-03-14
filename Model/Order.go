package Model

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	OID          uuid.UUID `gorm:"primary_key"`
	OTranID      string    `gorm:"not null"`
	OPrice       float64   `gorm:"not null"`
	ODestination string    `gorm:"not null"`
	OStatus      string    `gorm:"not null"`
	OPaid        bool      `gorm:"not null"`
	OCreated     time.Time `gorm:"not null"`
}
