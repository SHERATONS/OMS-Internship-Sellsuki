package Entities

type Stock struct {
	StockID   int64 `json:"stockID"`
	ProductID int64 `json:"productID"`
	Quantity  int64 `json:"quantity"`
}
