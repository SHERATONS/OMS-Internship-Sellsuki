package Order

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

func (o *Order) ValidateTranID(rawData map[string]interface{}) error {
	if TranID, ok := rawData["OTranID"].(string); ok {
		if reflect.TypeOf(TranID).Kind() != reflect.String {
			return errors.New("transaction ID Must Be a String")
		}
	} else {
		return errors.New("transaction ID is Required and Must Be a string")
	}

	return nil
}

func (o *Order) ValidateOrderStatus(rawData map[string]interface{}) error {
	if OStatus, ok := rawData["OStatus"].(string); ok {
		if reflect.TypeOf(OStatus).Kind() != reflect.String {
			return errors.New("order Status Must Be String")
		}
	} else {
		return errors.New("order Status is Required and Must Be String")
	}

	return nil
}

func (o *Order) ChangeStatus(order Order, orderStatus string) (Order, error) {
	switch orderStatus {
	case "Paid":
		if order.OStatus == "New" {
			order.OStatus = "Paid"
			order.OPaid = true
		} else {
			return order, errors.New("wrong Order Process")
		}

	case "Processing":
		if order.OStatus == "Paid" {
			if order.ODestination != "Branch" {
				order.OStatus = "Processing"
			} else {
				return order, errors.New("please Come Pick Up your Product at the Branch")
			}
		} else {
			return order, errors.New("wrong Order Process")
		}

	case "Done":
		if order.OStatus == "Processing" || (order.OStatus == "Paid" && order.ODestination == "Branch") {
			order.OStatus = "Done"
		} else {
			return order, errors.New("wrong Order Process")
		}

	default:
		return order, errors.New("invalid Order Status")
	}

	return order, nil
}
