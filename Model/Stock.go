package Model

import "time"

type Stock struct {
	SID       string    `gorm:"primary_key"`
	SQuantity int8      `gorm:"NOT NULL"`
	SUpdated  time.Time `gorm:"DEFAULT:current_timestamp"`
}
