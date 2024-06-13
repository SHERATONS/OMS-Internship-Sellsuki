package UseCases

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository"
)

type ProductUseCases struct {
	Repo Repository.IProductRepo
}

func (p ProductUseCases) UpdateProduct(product Entities.Product, productID string) (Entities.Product, error) {
	return p.Repo.UpdateProduct(product, productID)
}

func (p ProductUseCases) GetProductById(productId string) (Entities.Product, error) {
	return p.Repo.GetProductByID(productId)
}

func (p ProductUseCases) CreateProduct(product Entities.Product) (Entities.Product, error) {
	return p.Repo.CreateProduct(product)
}

func (p ProductUseCases) DeleteProductById(productId string) error {
	return p.Repo.DeleteProduct(productId)
}

func (p ProductUseCases) GetAllProducts() ([]Entities.Product, error) {
	return p.Repo.GetAllProducts()
}

func NewProductUseCases(Repo Repository.IProductRepo) IProductCase {
	return ProductUseCases{Repo: Repo}
}
