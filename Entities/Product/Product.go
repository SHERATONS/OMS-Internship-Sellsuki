package Product

import (
	"errors"

	"reflect"
	"strconv"
	"time"
)

type Product struct {
	PID      string
	PName    string
	PPrice   float64
	PDesc    string
	PCreated time.Time
	PUpdated time.Time
}

func (product *Product) ValidateProductID(rawData map[string]interface{}) error {
	if pId, ok := rawData["PID"]; ok {
		if reflect.TypeOf(pId).Kind() != reflect.String {
			return errors.New("product ID Must Be String")
		} else {
			CheckIdString := pId.(string)
			if CheckIdInt, err := strconv.Atoi(CheckIdString); err != nil {
				return errors.New("product ID Must Be a Number")
			} else if CheckIdInt <= 0 {
				return errors.New("product ID Must Be Greater than 0")
			}
		}
	} else {
		return errors.New("product ID is Required and Must Be String")
	}
	return nil
}

func (product *Product) ValidateProductName(rawData map[string]interface{}) error {
	if pName, ok := rawData["PName"]; ok {
		if reflect.TypeOf(pName).Kind() != reflect.String {
			return errors.New("product Name Must Be String")
		}
	} else {
		return errors.New("product Name is Required and Must Be String")
	}

	return nil
}

func (product *Product) ValidateProductPrice(rawData map[string]interface{}) error {
	if pPrice, ok := rawData["PPrice"]; ok {
		if reflect.TypeOf(pPrice).Kind() != reflect.Float64 {
			return errors.New("product Price Must Be Float")
		} else {
			CheckPriceFloat := pPrice.(float64)
			if CheckPriceFloat <= 0 {
				return errors.New("product Price Must Be Greater than 0")
			}
		}
	} else {
		return errors.New("product Price is Required and Must Be Float")
	}

	return nil
}

func (product *Product) ValidateProductDescription(rawData map[string]interface{}) error {
	if pDesc, ok := rawData["PDesc"]; ok {
		if reflect.TypeOf(pDesc).Kind() != reflect.String {
			return errors.New("product Description Must Be String")
		}
	} else {
		return errors.New("product Description is Required and Must Be String")
	}

	return nil
}
