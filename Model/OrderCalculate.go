package Model

type OrderCalculate struct {
	OTranID      string  `gorm:"primary_key"`
	OTotalPrice  float64 `gorm:"NOT NULL"`
	ODestination string  `gorm:"NOT NULL"`
	OProduct     string  `gorm:"NOT NULL"`
}
