package Model

import "time"

type Order struct {
	O_TranID     int64 `gorm:"primary_key"`
	O_CreateTime time.Time
	O_Status     int
	O_TotalPrice float64
	O_Delivery   bool
}
