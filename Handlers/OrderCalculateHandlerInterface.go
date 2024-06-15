package Handlers

import "github.com/gofiber/fiber/v2"

type OrderCalculateHandlerI interface {
	//GetAllOrder(c *fiber.Ctx) error
	//GetOrderByTransactionID(c *fiber.Ctx) error
	CreateTransactionID(c *fiber.Ctx) error
	//DeleteTransactionID(c *fiber.Ctx) error
}
