package Stock

import "github.com/gofiber/fiber/v2"

type IStockHandler interface {
	GetAllStock(c *fiber.Ctx) error
	GetStockByID(c *fiber.Ctx) error
	CreateStock(c *fiber.Ctx) error
	UpdateStock(c *fiber.Ctx) error
	DeleteStock(c *fiber.Ctx) error
}
