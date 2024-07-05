package Order

import (
	Order2 "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Order"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	UseCase UseCases.IOrderUseCase
}

func (o *OrderHandler) GetOrderById(c *fiber.Ctx) error {
	ctx, span := tracer.Start(c.UserContext(), "GetOrderById_Handler")
	defer span.End()

	orderID := c.Params("oid")

	order, err := o.UseCase.GetOrderById(ctx, orderID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Order Id Not Found"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"order": order})
}

func (o *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	ctx, span := tracer.Start(c.UserContext(), "CreateOrder_Handler")
	defer span.End()

	var rawData map[string]interface{}
	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	var validationError []string

	var tempOrder Order2.Order

	if err := tempOrder.ValidateTranID(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	order, err := o.UseCase.CreateOrder(ctx, rawData["OTranID"].(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Created Order Succeed",
		"orderID": order.OID})
}

func (o *OrderHandler) ChangeOrderStatus(c *fiber.Ctx) error {
	ctx, span := tracer.Start(c.UserContext(), "ChangeOrderStatus_Handler")
	defer span.End()

	var rawData map[string]interface{}

	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	var validationError []string

	var tempOrder Order2.Order

	if err := tempOrder.ValidateOrderStatus(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	orderStatus := rawData["OStatus"].(string)

	orderID := c.Params("oid")

	order, err := o.UseCase.ChangeOrderStatus(ctx, orderID, orderStatus)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"order": order})
}

func NewOrderHandler(useCase UseCases.IOrderUseCase) IOrderHandler {
	return &OrderHandler{UseCase: useCase}
}
