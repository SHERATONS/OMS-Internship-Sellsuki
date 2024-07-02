package Product

import (
	"testing"
)

var tempProduct Product

func TestValidateProductID(t *testing.T) {
	tests := []struct {
		rawData  map[string]interface{}
		expected string
	}{
		{map[string]interface{}{"PID": "123"}, ""},
		{map[string]interface{}{"PID": "abc"}, "product ID Must Be a Number"},
		{map[string]interface{}{"PID": -1}, "product ID Must Be String"},
		{map[string]interface{}{"PID": 123}, "product ID Must Be String"},
		{map[string]interface{}{"PID": "-1"}, "product ID Must Be Greater than 0"},
		{map[string]interface{}{}, "product ID is Required and Must Be String"},
	}

	for _, test := range tests {
		err := tempProduct.ValidateProductID(test.rawData)
		if err != nil && err.Error() != test.expected {
			t.Errorf("ValidateProductID(%v) = %v; want %v", test.rawData, err.Error(), test.expected)
		} else if err == nil && test.expected != "" {
			t.Errorf("ValidateProductID(%v) = nil; want %v", test.rawData, test.expected)
		}
	}
}

func TestValidateProductName(t *testing.T) {
	tests := []struct {
		rawData  map[string]interface{}
		expected string
	}{
		{map[string]interface{}{"PName": "Product1"}, ""},
		{map[string]interface{}{"PName": 123}, "product Name Must Be String"},
		{map[string]interface{}{}, "product Name is Required and Must Be String"},
	}

	for _, test := range tests {
		err := tempProduct.ValidateProductName(test.rawData)
		if err != nil && err.Error() != test.expected {
			t.Errorf("ValidateProductName(%v) = %v; want %v", test.rawData, err.Error(), test.expected)
		} else if err == nil && test.expected != "" {
			t.Errorf("ValidateProductName(%v) = nil; want %v", test.rawData, test.expected)
		}
	}
}

func TestValidateProductPrice(t *testing.T) {
	tests := []struct {
		rawData  map[string]interface{}
		expected string
	}{
		{map[string]interface{}{"PPrice": 99.99}, ""},
		{map[string]interface{}{"PPrice": "abc"}, "product Price Must Be Float"},
		{map[string]interface{}{"PPrice": -99.99}, "product Price Must Be Greater than 0"},
		{map[string]interface{}{}, "product Price is Required and Must Be Float"},
	}

	for _, test := range tests {
		err := tempProduct.ValidateProductPrice(test.rawData)
		if err != nil && err.Error() != test.expected {
			t.Errorf("ValidateProductPrice(%v) = %v; want %v", test.rawData, err.Error(), test.expected)
		} else if err == nil && test.expected != "" {
			t.Errorf("ValidateProductPrice(%v) = nil; want %v", test.rawData, test.expected)
		}
	}
}

func TestValidateProductDescription(t *testing.T) {
	tests := []struct {
		rawData  map[string]interface{}
		expected string
	}{
		{map[string]interface{}{"PDesc": "A great product"}, ""},
		{map[string]interface{}{"PDesc": 123}, "product Description Must Be String"},
		{map[string]interface{}{}, "product Description is Required and Must Be String"},
	}

	for _, test := range tests {
		err := tempProduct.ValidateProductDescription(test.rawData)
		if err != nil && err.Error() != test.expected {
			t.Errorf("ValidateProductDescription(%v) = %v; want %v", test.rawData, err.Error(), test.expected)
		} else if err == nil && test.expected != "" {
			t.Errorf("ValidateProductDescription(%v) = nil; want %v", test.rawData, test.expected)
		}
	}
}
