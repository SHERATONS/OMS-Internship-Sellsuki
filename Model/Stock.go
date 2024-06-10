package Model

type Stock struct {
	SID       string `gorm:"primary_key;AUTO_INCREMENT"`
	SQuantity int64
}
