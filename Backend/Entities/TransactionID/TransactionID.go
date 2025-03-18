package TransactionID

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"reflect"
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

func (order *TransactionID) ValidateTDestination(rawData map[string]interface{}) error {
	if destination, ok := rawData["TDestination"]; ok {
		if reflect.TypeOf(destination).Kind() != reflect.String {
			return errors.New("destination Must Be a String")
		}
	} else {
		return errors.New("destination is Required and Must Be a string")
	}

	return nil
}

func (order *TransactionID) ValidateProductList(rawData map[string]interface{}) error {
	if product, ok := rawData["TProductList"]; ok {
		if reflect.TypeOf(product).Kind() != reflect.String {
			return errors.New("product Must Be a String")
		}
	} else {
		return errors.New("product is Required and Must Be a string")
	}

	return nil
}
