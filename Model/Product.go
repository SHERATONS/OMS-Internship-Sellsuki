package Model

type Product struct {
	P_ID    int `gorm:"primary_key;AUTO_INCREMENT"`
	P_Name  string
	P_Price float64
	P_Desc  string
}
