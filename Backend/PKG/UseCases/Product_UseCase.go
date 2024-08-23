package UseCases

import (
	"context"
	Product2 "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Product"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/PKG/Repository/Product"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/PKG/Repository/Stock"
)

type ProductUseCase struct {
	Repo      Product.IProductRepo
	RepoStock Stock.IStockRepo
}

func (p ProductUseCase) UpdateProduct(ctx context.Context, product Product2.Product, productID string) (Product2.Product, error) {
	ctx, span := tracerProduct.Start(ctx, "UpdateProduct_UseCase")
	defer span.End()

	tempProduct, err := p.Repo.GetProductByID(ctx, productID)
	if err != nil {
		return Product2.Product{}, err
	}

	return p.Repo.UpdateProduct(ctx, product, productID, tempProduct)
}

func (p ProductUseCase) GetProductById(ctx context.Context, productID string) (Product2.Product, error) {
	ctx, span := tracerProduct.Start(ctx, "GetProductById_UseCase")
	defer span.End()

	return p.Repo.GetProductByID(ctx, productID)
}

func (p ProductUseCase) CreateProduct(ctx context.Context, product Product2.Product) (Product2.Product, error) {
	ctx, span := tracerProduct.Start(ctx, "CreateProduct_UseCase")
	defer span.End()

	return p.Repo.CreateProduct(ctx, product)
}

func (p ProductUseCase) DeleteProductById(ctx context.Context, productID string) error {
	ctx, span := tracerProduct.Start(ctx, "DeleteProductById_UseCase")
	defer span.End()

	_, err := p.Repo.GetProductByID(ctx, productID)
	if err != nil {
		return err
	}

	_, err = p.RepoStock.GetStockByID(ctx, productID)
	if err == nil {
		err = p.RepoStock.DeleteStock(ctx, productID)
	}

	return p.Repo.DeleteProduct(ctx, productID)
}

func (p ProductUseCase) GetAllProducts(ctx context.Context) ([]Product2.Product, error) {
	ctx, span := tracerProduct.Start(ctx, "GetAllProducts_UseCase")
	defer span.End()

	return p.Repo.GetAllProducts(ctx)
}

func NewProductUseCase(Repo Product.IProductRepo, RepoStock Stock.IStockRepo) IProductUseCase {
	return ProductUseCase{Repo: Repo, RepoStock: RepoStock}
}
