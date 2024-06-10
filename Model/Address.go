package Model

type Address struct {
	AID     string `gorm:"primary_key;AUTO_INCREMENT"`
	City    string
	Country string
}
