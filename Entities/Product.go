package Entities

type product struct {
	ProductID    int64  `json:"productId"`
	ProductName  string `json:"productName"`
	ProductPrice int64  `json:"productPrice"`
	ProductDesc  string `json:"productDesc"`
}
