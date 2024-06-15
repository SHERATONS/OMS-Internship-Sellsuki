package Handlers

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
	"github.com/gofiber/fiber/v2"
)

type OrderCalculateHandler struct {
	UseCases        UseCases.IOrderCalculateCase
	UseCasesProduct UseCases.IProductCase
	UseCasesAddress UseCases.IAddressCase
}

//func (o *OrderCalculateHandler) GetAllOrder(c *fiber.Ctx) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (o *OrderCalculateHandler) GetOrderByTransactionID(c *fiber.Ctx) error {
//	//TODO implement me
//	panic("implement me")
//}

func (o *OrderCalculateHandler) CreateTransactionID(c *fiber.Ctx) error {
	var rawData map[string]interface{}
	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Request Body")
	}
	
}

//func (o *OrderCalculateHandler) DeleteTransactionID(c *fiber.Ctx) error {
//	//TODO implement me
//	panic("implement me")
//}

func NewOrderCalculateHandler(useCases UseCases.IOrderCalculateCase, useCasesProduct UseCases.IProductCase, useCasesAddress UseCases.IAddressCase) OrderCalculateHandlerI {
	return &OrderCalculateHandler{
		UseCases:        useCases,
		UseCasesProduct: useCasesProduct,
		UseCasesAddress: useCasesAddress,
	}
}
