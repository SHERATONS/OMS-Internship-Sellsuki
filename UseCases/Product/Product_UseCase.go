package Product

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository/Product"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository/Stock"
	"time"
)

type ProductUseCase struct {
	Repo      Product.IProductRepo
	RepoStock Stock.IStockRepo
}

func (p ProductUseCase) UpdateProduct(product Entities.Product, productID string) (Entities.Product, error) {
	tempProduct, err := p.Repo.GetProductByID(productID)
	if err != nil {
		return product, err
	}

	product.PCreated = tempProduct.PCreated
	product.PUpdated = time.Now()

	return p.Repo.UpdateProduct(product, productID)
}

func (p ProductUseCase) GetProductById(productID string) (Entities.Product, error) {
	return p.Repo.GetProductByID(productID)
}

func (p ProductUseCase) CreateProduct(product Entities.Product) (Entities.Product, error) {
	product.PCreated = time.Now()
	product.PUpdated = time.Now()

	return p.Repo.CreateProduct(product)
}

func (p ProductUseCase) DeleteProductById(productID string) error {
	_, err := p.Repo.GetProductByID(productID)
	if err != nil {
		return err
	}

	_, err = p.RepoStock.GetStockByID(productID)
	if err == nil {
		err = p.RepoStock.DeleteStock(productID)
	}

	return p.Repo.DeleteProduct(productID)
}

func (p ProductUseCase) GetAllProducts() ([]Entities.Product, error) {
	return p.Repo.GetAllProducts()
}

func NewProductUseCase(Repo Product.IProductRepo, RepoStock Stock.IStockRepo) IProductUseCase {
	return ProductUseCase{Repo: Repo, RepoStock: RepoStock}
}
