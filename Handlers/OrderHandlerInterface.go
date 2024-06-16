package Handlers

import (
	"github.com/gofiber/fiber/v2"
)

type OrderHandlerI interface {
	CreateOrder(c *fiber.Ctx) error
	ChangeOrderStatus(c *fiber.Ctx) error
	GetOrderById(c *fiber.Ctx) error
}
