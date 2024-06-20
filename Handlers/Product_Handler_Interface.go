package Handlers

import "github.com/gofiber/fiber/v2"

type IProductHandler interface {
	GetAllProducts(c *fiber.Ctx) error
	GetProductByID(c *fiber.Ctx) error
	CreateProduct(c *fiber.Ctx) error
	UpdateProductById(c *fiber.Ctx) error
	DeleteProductById(c *fiber.Ctx) error
}
