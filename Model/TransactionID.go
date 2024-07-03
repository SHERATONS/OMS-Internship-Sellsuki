package Model

type TransactionID struct {
	TID          string  `gorm:"primary_key"`
	TPrice       float64 `gorm:"NOT NULL"`
	TDestination string  `gorm:"NOT NULL"`
	TProductList string  `gorm:"NOT NULL"`
}
