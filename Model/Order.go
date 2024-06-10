package Model

import "time"

type Order struct {
	OTranID     string `gorm:"primary_key"`
	OCreateTime time.Time
	OStatus     string
	OTotalPrice float64
	ODelivery   bool
}
