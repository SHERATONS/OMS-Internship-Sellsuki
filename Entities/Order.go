package Entities

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	OID          uuid.UUID
	OTranID      string
	OPrice       float64
	ODestination string
	OStatus      string
	OPaid        bool
	OCreated     time.Time
}
