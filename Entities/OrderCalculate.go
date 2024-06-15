package Entities

import (
	"crypto/sha1"
	"fmt"
	"time"
)

type OrderCalculate struct {
	OTranID      string
	OTotalPrice  float64
	ODestination string
	OProduct     string
}

func (order *OrderCalculate) Tablename() string {
	return "ordercalculates"
}

func (order *OrderCalculate) GenerateTransactionID(orderPrice float64) string {
	currentTime := time.Now().Format("20060102150405") // Format: YYYYMMDDHHMMSS
	data := fmt.Sprintf("%f%s", orderPrice, currentTime)
	hash := sha1.New()
	hash.Write([]byte(data))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
