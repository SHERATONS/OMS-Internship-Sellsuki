package TransactionID

import (
	"testing"
)

var tempTransactionID TransactionID

func TestValidateTDestination(t *testing.T) {
	tests := []struct {
		rawData  map[string]interface{}
		expected string
	}{
		{map[string]interface{}{"TDestination": "Branch"}, ""},
		{map[string]interface{}{"TDestination": 123}, "destination Must Be a String"},
		{map[string]interface{}{}, "destination is Required and Must Be a string"},
	}

	for _, test := range tests {
		err := tempTransactionID.ValidateTDestination(test.rawData)
		if err != nil && err.Error() != test.expected {
			t.Errorf("ValidateTDestination(%v) = %v; want %v", test.rawData, err.Error(), test.expected)
		} else if err == nil && test.expected != "" {
			t.Errorf("ValidateTDestination(%v) = nil; want %v", test.rawData, test.expected)
		}
	}
}

func TestValidateProductList(t *testing.T) {
	tests := []struct {
		rawData  map[string]interface{}
		expected string
	}{
		{map[string]interface{}{"TProductList": "Product1:Q1, Product2:Q2"}, ""},
		{map[string]interface{}{"TProductList": 123}, "product Must Be a String"},
		{map[string]interface{}{}, "product is Required and Must Be a string"},
	}

	for _, test := range tests {
		err := tempTransactionID.ValidateProductList(test.rawData)
		if err != nil && err.Error() != test.expected {
			t.Errorf("ValidateProductList(%v) = %v; want %v", test.rawData, err.Error(), test.expected)
		} else if err == nil && test.expected != "" {
			t.Errorf("ValidateProductList(%v) = nil; want %v", test.rawData, test.expected)
		}
	}
}
