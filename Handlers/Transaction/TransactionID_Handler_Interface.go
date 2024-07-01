package Transaction

import "github.com/gofiber/fiber/v2"

type ITransactionIDHandler interface {
	GetAllTransactionIDs(c *fiber.Ctx) error
	GetOrderByTransactionID(c *fiber.Ctx) error
	CreateTransactionID(c *fiber.Ctx) error
	DeleteTransactionID(c *fiber.Ctx) error
}
