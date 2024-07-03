package Order

import (
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
)

type IOrderHandler interface {
	CreateOrder(c *fiber.Ctx) error
	ChangeOrderStatus(c *fiber.Ctx) error
	GetOrderById(c *fiber.Ctx) error
}

var tracer = otel.Tracer("Order_")
