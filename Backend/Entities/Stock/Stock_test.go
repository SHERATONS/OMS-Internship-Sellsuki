package Stock

import (
	"testing"
)

var tempStock Stock

func TestValidateStockID(t *testing.T) {
	tests := []struct {
		rawData  map[string]interface{}
		expected string
	}{
		{map[string]interface{}{"SID": "123"}, ""},
		{map[string]interface{}{"SID": 123}, "stock ID Must Be a String"},
		{map[string]interface{}{"SID": "abc"}, "stock ID Must Be a Number"},
		{map[string]interface{}{"SID": "-123"}, "stock ID Must Be Greater than 0"},
		{map[string]interface{}{}, "stock ID is Required and Must Be a String"},
	}

	for _, test := range tests {
		err := tempStock.ValidateStockID(test.rawData)
		if err != nil && err.Error() != test.expected {
			t.Errorf("ValidateStockID(%v) = %v; want %v", test.rawData, err.Error(), test.expected)
		} else if err == nil && test.expected != "" {
			t.Errorf("ValidateStockID(%v) = nil; want %v", test.rawData, test.expected)
		}
	}
}

func TestValidateStockQuantity(t *testing.T) {
	tests := []struct {
		rawData  map[string]interface{}
		expected string
	}{
		{map[string]interface{}{"SQuantity": 10.0}, ""},
		{map[string]interface{}{"SQuantity": "10"}, "stock Quantity Must Be a Integer"},
		{map[string]interface{}{"SQuantity": -10.0}, "stock Quantity Must Be Greater than 0"},
		{map[string]interface{}{}, "stock Quantity is Required and Must Be a Integer"},
	}

	for _, test := range tests {
		err := tempStock.ValidateStockQuantity(test.rawData)
		if err != nil && err.Error() != test.expected {
			t.Errorf("ValidateStockQuantity(%v) = %v; want %v", test.rawData, err.Error(), test.expected)
		} else if err == nil && test.expected != "" {
			t.Errorf("ValidateStockQuantity(%v) = nil; want %v", test.rawData, test.expected)
		}
	}
}
