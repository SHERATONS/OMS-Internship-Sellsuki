package Entities

import (
	"crypto/sha1"
	"fmt"
	"time"
)

type TransactionID struct {
	TID          string
	TPrice       float64
	TDestination string
	TProductList string
}

func (order *TransactionID) GenerateTransactionID(orderPrice float64) string {
	currentTime := time.Now().Format("20060102150405") // Format: YYYYMMDDHHMMSS
	data := fmt.Sprintf("%f%s", orderPrice, currentTime)
	hash := sha1.New()
	hash.Write([]byte(data))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
