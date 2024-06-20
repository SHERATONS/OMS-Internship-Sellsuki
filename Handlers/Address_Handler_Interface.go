package Handlers

import (
	"github.com/gofiber/fiber/v2"
)

type IAddressHandler interface {
	GetAddressByCity(c *fiber.Ctx) error
	CreateAddress(c *fiber.Ctx) error
	UpdateAddress(c *fiber.Ctx) error
	DeleteAddress(c *fiber.Ctx) error
}
