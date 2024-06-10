package Model

type Product struct {
	PID    string `gorm:"primary_key;AUTO_INCREMENT"`
	PName  string
	PPrice float64
	PDesc  string
}
