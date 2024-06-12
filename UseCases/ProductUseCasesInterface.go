package UseCases

import "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"

type IProductCase interface {
	GetAllProducts() ([]Entities.Product, error)
	GetProductById(productId string) (Entities.Product, error)
	CreateProduct(product Entities.Product) (Entities.Product, error)
	UpdateProduct(product Entities.Product, productId string) (Entities.Product, error)
	DeleteProductById(productId string) error
}
