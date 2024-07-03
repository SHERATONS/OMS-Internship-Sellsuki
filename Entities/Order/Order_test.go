package Order

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

var tempOrder Order

func TestValidateTranID(t *testing.T) {
	tests := []struct {
		rawData  map[string]interface{}
		expected string
	}{
		{map[string]interface{}{"OTranID": "12345"}, ""},
		{map[string]interface{}{"OTranID": 12345}, "transaction ID is Required and Must Be a string"},
		{map[string]interface{}{}, "transaction ID is Required and Must Be a string"},
	}

	for _, test := range tests {
		err := tempOrder.ValidateTranID(test.rawData)
		if err != nil && err.Error() != test.expected {
			t.Errorf("ValidateTranID(%v) = %v; want %v", test.rawData, err.Error(), test.expected)
		} else if err == nil && test.expected != "" {
			t.Errorf("ValidateTranID(%v) = nil; want %v", test.rawData, test.expected)
		}
	}
}

func TestValidateOrderStatus(t *testing.T) {
	tests := []struct {
		rawData  map[string]interface{}
		expected string
	}{
		{map[string]interface{}{"OStatus": "New"}, ""},
		{map[string]interface{}{"OStatus": 123}, "order Status is Required and Must Be String"},
		{map[string]interface{}{}, "order Status is Required and Must Be String"},
	}

	for _, test := range tests {
		err := tempOrder.ValidateOrderStatus(test.rawData)
		if err != nil && err.Error() != test.expected {
			t.Errorf("ValidateOrderStatus(%v) = %v; want %v", test.rawData, err.Error(), test.expected)
		} else if err == nil && test.expected != "" {
			t.Errorf("ValidateOrderStatus(%v) = nil; want %v", test.rawData, test.expected)
		}
	}
}

func TestChangeStatus(t *testing.T) {
	tests := []struct {
		initialStatus string
		newStatus     string
		destination   string
		expected      string
	}{
		{"New", "Paid", "", ""},
		{"Paid", "Processing", "Bangkok", ""},
		{"Processing", "Done", "", ""},
		{"Paid", "Done", "Branch", ""},
		{"New", "Processing", "", "wrong Order Process"},
		{"New", "Done", "", "wrong Order Process"},
		{"Paid", "Processing", "Branch", "please Come Pick Up your Product at the Branch"},
	}

	for _, test := range tests {
		order := Order{
			OID:          uuid.UUID{},
			OTranID:      "12345",
			OPrice:       100.0,
			ODestination: test.destination,
			OStatus:      test.initialStatus,
			OPaid:        true,
			OCreated:     time.Now(),
		}
		_, err := order.ChangeStatus(order, test.newStatus)
		if err != nil && err.Error() != test.expected {
			t.Errorf("ChangeStatus(%v, %v) = %v; want %v", test.initialStatus, test.newStatus, err.Error(), test.expected)
		} else if err == nil && test.expected != "" {
			t.Errorf("ChangeStatus(%v, %v) = nil; want %v", test.initialStatus, test.newStatus, test.expected)
		}
	}
}
