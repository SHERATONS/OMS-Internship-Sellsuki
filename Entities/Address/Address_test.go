package Address

import (
	"testing"
)

var tempAddress Address

func TestValidateCity(t *testing.T) {
	tests := []struct {
		rawData  map[string]interface{}
		expected string
	}{
		{map[string]interface{}{"City": "Bangkok"}, ""},
		{map[string]interface{}{"City": 123}, "city Must Be a String"},
		{map[string]interface{}{}, "city is Required and Must Be a String"},
	}

	for _, test := range tests {
		err := tempAddress.ValidateCity(test.rawData)
		if err != nil && err.Error() != test.expected {
			t.Errorf("ValidateCity(%v) = %v; want %v", test.rawData, err.Error(), test.expected)
		} else if err == nil && test.expected != "" {
			t.Errorf("ValidateCity(%v) = nil; want %v", test.rawData, test.expected)
		}
	}
}

func TestValidateCountry(t *testing.T) {
	tests := []struct {
		rawData  map[string]interface{}
		expected string
	}{
		{map[string]interface{}{"Country": "Thailand"}, ""},
		{map[string]interface{}{"Country": 123}, "country Must Be a String"},
		{map[string]interface{}{}, "country is Required and Must Be a String"},
	}

	for _, test := range tests {
		err := tempAddress.ValidateCountry(test.rawData)
		if err != nil && err.Error() != test.expected {
			t.Errorf("ValidateCountry(%v) = %v; want %v", test.rawData, err.Error(), test.expected)
		} else if err == nil && test.expected != "" {
			t.Errorf("ValidateCountry(%v) = nil; want %v", test.rawData, test.expected)
		}
	}
}

func TestValidateAPrice(t *testing.T) {
	tests := []struct {
		rawData  map[string]interface{}
		expected string
	}{
		{map[string]interface{}{"APrice": 10.0}, ""},
		{map[string]interface{}{"APrice": "10"}, "price Must Be a Float"},
		{map[string]interface{}{"APrice": -10.0}, "price Must Be Greater than 0"},
		{map[string]interface{}{}, "price is Required and Must Be a Float"},
	}

	for _, test := range tests {
		err := tempAddress.ValidateAPrice(test.rawData)
		if err != nil && err.Error() != test.expected {
			t.Errorf("ValidateAPrice(%v) = %v; want %v", test.rawData, err.Error(), test.expected)
		} else if err == nil && test.expected != "" {
			t.Errorf("ValidateAPrice(%v) = nil; want %v", test.rawData, test.expected)
		}
	}
}
