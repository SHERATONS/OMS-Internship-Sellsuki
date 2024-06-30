package Entities

import (
	"errors"
	"github.com/google/uuid"
	"reflect"
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

func ValidateTranID(rawData map[string]interface{}) error {
	if TranID, ok := rawData["OTranID"].(string); ok {
		if reflect.TypeOf(TranID).Kind() != reflect.String {
			return errors.New("transaction ID Must Be a String")
		}
	} else {
		return errors.New("transaction ID is Required and Must Be a string")
	}

	return nil
}

func ValidateOrderStatus(rawData map[string]interface{}) error {
	if OStatus, ok := rawData["OStatus"].(string); ok {
		if reflect.TypeOf(OStatus).Kind() != reflect.String {
			return errors.New("order Status Must Be String")
		}
	} else {
		return errors.New("order Status is Required and Must Be String")
	}

	return nil
}
