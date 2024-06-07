package Model

type Address struct {
	A_ID    int64 `gorm:"primary_key;AUTO_INCREMENT"`
	City    string
	Country string
}
