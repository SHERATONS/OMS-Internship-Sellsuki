package UseCases

import (
	"context"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Product"
	"go.opentelemetry.io/otel"
)

type IProductUseCase interface {
	GetAllProducts(ctx context.Context) ([]Product.Product, error)
	GetProductById(ctx context.Context, productId string) (Product.Product, error)
	CreateProduct(ctx context.Context, product Product.Product) (Product.Product, error)
	UpdateProduct(ctx context.Context, product Product.Product, productId string) (Product.Product, error)
	DeleteProductById(ctx context.Context, productId string) error
}

var tracerProduct = otel.Tracer("Product_UseCase")
