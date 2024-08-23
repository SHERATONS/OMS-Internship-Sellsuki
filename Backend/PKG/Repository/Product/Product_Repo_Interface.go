package Product

import (
	"context"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Product"
	"go.opentelemetry.io/otel"
)

type IProductRepo interface {
	GetAllProducts(ctx context.Context) ([]Product.Product, error)
	GetProductByID(ctx context.Context, productId string) (Product.Product, error)
	CreateProduct(ctx context.Context, product Product.Product) (Product.Product, error)
	UpdateProduct(ctx context.Context, product Product.Product, productID string, tempProduct Product.Product) (Product.Product, error)
	DeleteProduct(ctx context.Context, productID string) error
}

var tracer = otel.Tracer("Product_Repo")
