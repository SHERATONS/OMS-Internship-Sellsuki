package Product

import (
	"context"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
)

type IProductRepo interface {
	GetAllProducts(ctx context.Context) ([]Entities.Product, error)
	GetProductByID(ctx context.Context, productId string) (Entities.Product, error)
	CreateProduct(ctx context.Context, product Entities.Product) (Entities.Product, error)
	UpdateProduct(ctx context.Context, product Entities.Product, productID string) (Entities.Product, error)
	DeleteProduct(ctx context.Context, productID string) error
}
