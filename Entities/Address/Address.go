package Address

import (
	"errors"
	"reflect"
	"time"
)

type Address struct {
	City     string
	Country  string
	APrice   float64
	AUpdated time.Time
}

func (address *Address) ValidateCity(rawData map[string]interface{}) error {
	if city, ok := rawData["City"]; ok {
		if reflect.TypeOf(city).Kind() != reflect.String {
			return errors.New("city Must Be a String")
		}
	} else {
		return errors.New("city is Required and Must Be a String")
	}

	return nil
}

func (address *Address) ValidateCountry(rawData map[string]interface{}) error {
	if country, ok := rawData["Country"]; ok {
		if reflect.TypeOf(country).Kind() != reflect.String {
			return errors.New("country Must Be a String")
		}
	} else {
		return errors.New("country is Required and Must Be a String")
	}

	return nil
}

func (address *Address) ValidateAPrice(rawData map[string]interface{}) error {
	if aPrice, ok := rawData["APrice"]; ok {
		if reflect.TypeOf(aPrice).Kind() != reflect.Float64 {
			return errors.New("price Must Be a Float")
		} else if CheckPriceFloat := aPrice.(float64); CheckPriceFloat <= 0 {
			return errors.New("price Must Be Greater than 0")
		}
	} else {
		return errors.New("price is Required and Must Be a Float")
	}

	return nil
}
