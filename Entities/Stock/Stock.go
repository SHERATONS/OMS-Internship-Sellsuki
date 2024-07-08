package Stock

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

type Stock struct {
	SID       string
	SQuantity float64
	SUpdated  time.Time
}

func (stock *Stock) ValidateStockID(rawData map[string]interface{}) error {
	if sId, ok := rawData["SID"]; ok {
		if reflect.TypeOf(sId).Kind() != reflect.String {
			return errors.New("stock ID Must Be a String")
		} else {
			checkIdString := sId.(string)
			if checkIdInt, err := strconv.Atoi(checkIdString); err != nil {
				return errors.New("stock ID Must Be a Number")
			} else if checkIdInt <= 0 {
				return errors.New("stock ID Must Be Greater than 0")
			}
		}
	} else {
		return errors.New("stock ID is Required and Must Be a String")
	}

	return nil
}

func (stock *Stock) ValidateStockQuantity(rawData map[string]interface{}) error {
	if sQuantity, ok := rawData["SQuantity"]; ok {
		if reflect.TypeOf(sQuantity).Kind() != reflect.Float64 {
			return errors.New("stock Quantity Must Be a Integer")
		} else {
			checkQuantityInt := sQuantity.(float64)
			if checkQuantityInt <= 0 {
				return errors.New("stock Quantity Must Be Greater than 0")
			}
		}
	} else {
		return errors.New("stock Quantity is Required and Must Be a Integer")
	}

	return nil
}
