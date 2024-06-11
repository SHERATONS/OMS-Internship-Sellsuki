package Entities

import "time"

type Product struct {
	ProductID         string
	ProductName       string
	ProductPrice      int64
	ProductDesc       string
	ProductCreateTime time.Time
	ProductUpdateTime time.Time
}
