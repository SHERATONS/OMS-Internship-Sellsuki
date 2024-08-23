package Product

import (
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
)

type IProductHandler interface {
	GetAllProducts(c *fiber.Ctx) error
	GetProductByID(c *fiber.Ctx) error
	CreateProduct(c *fiber.Ctx) error
	UpdateProductById(c *fiber.Ctx) error
	DeleteProductById(c *fiber.Ctx) error
}

var tracer = otel.Tracer("Product_Handler")
