package Handlers

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	UseCase              UseCases.IOrderUseCase
	UseCaseStock         UseCases.IStockUseCase
	UseCaseTransactionID UseCases.ITransactionIDUseCase
}

func (o *OrderHandler) GetOrderById(c *fiber.Ctx) error {
	orderID := c.Params("oid")

	order, err := o.UseCase.GetOrderById(orderID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Order Id Not Found"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"order": order})
}

func (o *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var rawData map[string]interface{}
	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	var validationError []string

	if err := Entities.ValidateTranID(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	_, err := o.UseCase.CreateOrder(rawData["OTranID"].(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Created Order Succeed"})
}

func (o *OrderHandler) ChangeOrderStatus(c *fiber.Ctx) error {
	var rawData map[string]interface{}

	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	var validationError []string

	if err := Entities.ValidateOrderStatus(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	OrderStatus := rawData["OStatus"].(string)

	OrderID := c.Params("oid")

	order, err := o.UseCase.ChangeOrderStatus(OrderID, OrderStatus)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"order": order})
}

func NewOrderHandler(useCase UseCases.IOrderUseCase, useCaseStock UseCases.IStockUseCase, useCaseTransactionID UseCases.ITransactionIDUseCase) IOrderHandler {
	return &OrderHandler{
		UseCase:              useCase,
		UseCaseStock:         useCaseStock,
		UseCaseTransactionID: useCaseTransactionID,
	}
}
