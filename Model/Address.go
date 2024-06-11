package Model

type Address struct {
	AID     string `gorm:"primary_key"`
	City    string `gorm:"NOT NULL"`
	Country string `gorm:"NOT NULL"`
}
