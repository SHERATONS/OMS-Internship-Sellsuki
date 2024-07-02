package Transaction

import (
	"encoding/json"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/TransactionID"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases/Transaction"
	"github.com/gofiber/fiber/v2"
)

type TransactionIDHandler struct {
	UseCase Transaction.ITransactionIDUseCase
}

func (o *TransactionIDHandler) GetAllTransactionIDs(c *fiber.Ctx) error {
	ctx, span := tracer.Start(c.UserContext(), "GetAllTransactionIDs_Handler")
	defer span.End()

	transactionID, err := o.UseCase.GetAllTransactionIDs(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": "Something went wrong"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"transaction_id": transactionID})
}

func (o *TransactionIDHandler) GetOrderByTransactionID(c *fiber.Ctx) error {
	ctx, span := tracer.Start(c.UserContext(), "GetOrderByTransactionID_Handler")
	defer span.End()

	transactionID := c.Params("tid")

	order, err := o.UseCase.GetOrderByTransactionID(ctx, transactionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Transaction ID Not Found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"TransactionID": order})
}

func (o *TransactionIDHandler) CreateTransactionID(c *fiber.Ctx) error {
	ctx, span := tracer.Start(c.UserContext(), "CreateTransactionID_Handler")
	defer span.End()

	var rawData map[string]interface{}

	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	var validationError []string

	var tempTransactionID TransactionID.TransactionID

	if err := tempTransactionID.ValidateTDestination(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := tempTransactionID.ValidateProductList(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	var createTransactionID TransactionID.TransactionID

	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}
	if err := json.Unmarshal(data, &createTransactionID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}

	transactionID, err := o.UseCase.CreateTransactionID(ctx, createTransactionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"TransactionID": transactionID})
}

func (o *TransactionIDHandler) DeleteTransactionID(c *fiber.Ctx) error {
	ctx, span := tracer.Start(c.UserContext(), "DeleteTransactionID_Handler")
	defer span.End()

	transactionID := c.Params("tid")

	err := o.UseCase.DeleteTransactionID(ctx, transactionID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":        "Transaction Id Deleted Successfully",
		"transaction_id": transactionID,
	})
}

func NewTransactionIDHandler(useCase Transaction.ITransactionIDUseCase) ITransactionIDHandler {
	return &TransactionIDHandler{UseCase: useCase}
}
