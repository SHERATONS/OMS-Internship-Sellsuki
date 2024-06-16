package Handlers

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	UseCases      UseCases.IOrderCase
	UseCasesStock UseCases.IStockCase
}

func (o *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	 
}

func (o *OrderHandler) ChangeOrderStatus(c *fiber.Ctx) error {

}

func NewOrderHandler(useCases UseCases.IOrderCase, useCasesStock UseCases.IStockCase) OrderHandlerI {
	return &OrderHandler{
		UseCases:      useCases,
		UseCasesStock: useCasesStock,
	}
}
