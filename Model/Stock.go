package Model

type Stock struct {
	S_ID       int64 `gorm:"primary_key;AUTO_INCREMENT"`
	P_ID       int64 `gorm:"primary_key;AUTO_INCREMENT"`
	S_Quantity int64
}
