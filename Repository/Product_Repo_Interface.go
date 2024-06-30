package Repository

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
)

type IProductRepo interface {
	GetAllProducts() ([]Entities.Product, error)
	GetProductByID(productId string) (Entities.Product, error)
	CreateProduct(product Entities.Product) (Entities.Product, error)
	UpdateProduct(product Entities.Product, productID string) (Entities.Product, error)
	DeleteProduct(productID string) error
}
