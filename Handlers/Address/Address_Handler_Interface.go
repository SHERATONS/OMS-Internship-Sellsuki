package Address

import (
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
)

type IAddressHandler interface {
	GetAddressByCity(c *fiber.Ctx) error
	CreateAddress(c *fiber.Ctx) error
	UpdateAddress(c *fiber.Ctx) error
	DeleteAddress(c *fiber.Ctx) error
}

var tracer = otel.Tracer("Address_Handler")
