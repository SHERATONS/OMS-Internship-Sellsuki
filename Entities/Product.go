package Entities

type Product struct {
	ProductID    int64  `json:"productId"`
	ProductName  string `json:"productName"`
	ProductPrice int64  `json:"productPrice"`
	ProductDesc  string `json:"productDesc"`
}
