package Entities

type OrderCalculate struct {
	OTranID      string
	OTotalPrice  float64
	ODestination string
	OProduct     string
}

func (order *OrderCalculate) Tablename() string {
	return "ordercalculates"
}
