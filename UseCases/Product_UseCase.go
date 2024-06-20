package UseCases

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository"
)

type ProductUseCase struct {
	Repo Repository.IProductRepo
}

func (p ProductUseCase) UpdateProduct(product Entities.Product, productID string) (Entities.Product, error) {
	return p.Repo.UpdateProduct(product, productID)
}

func (p ProductUseCase) GetProductById(productID string) (Entities.Product, error) {
	return p.Repo.GetProductByID(productID)
}

func (p ProductUseCase) CreateProduct(product Entities.Product) (Entities.Product, error) {
	return p.Repo.CreateProduct(product)
}

func (p ProductUseCase) DeleteProductById(productID string) error {
	return p.Repo.DeleteProduct(productID)
}

func (p ProductUseCase) GetAllProducts() ([]Entities.Product, error) {
	return p.Repo.GetAllProducts()
}

func NewProductUseCase(Repo Repository.IProductRepo) IProductUseCase {
	return ProductUseCase{Repo: Repo}
}
