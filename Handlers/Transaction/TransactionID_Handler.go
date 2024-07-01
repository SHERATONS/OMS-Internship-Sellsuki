package Transaction

import (
	"encoding/json"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases/Transaction"
	"github.com/gofiber/fiber/v2"
)

type TransactionIDHandler struct {
	UseCase Transaction.ITransactionIDUseCase
}

func (o *TransactionIDHandler) GetAllTransactionIDs(c *fiber.Ctx) error {
	transactionID, err := o.UseCase.GetAllTransactionIDs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": "Something went wrong"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"transaction_id": transactionID})
}

func (o *TransactionIDHandler) GetOrderByTransactionID(c *fiber.Ctx) error {
	transactionID := c.Params("tid")

	order, err := o.UseCase.GetOrderByTransactionID(transactionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Transaction ID Not Found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"TransactionID": order})
}

func (o *TransactionIDHandler) CreateTransactionID(c *fiber.Ctx) error {
	var rawData map[string]interface{}

	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	var validationError []string

	if err := Entities.ValidateTDestination(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := Entities.ValidateProductList(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	var createTransactionID Entities.TransactionID

	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}
	if err := json.Unmarshal(data, &createTransactionID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}

	transactionID, err := o.UseCase.CreateTransactionID(createTransactionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"TransactionID": transactionID})
}

func (o *TransactionIDHandler) DeleteTransactionID(c *fiber.Ctx) error {
	transactionID := c.Params("tid")

	err := o.UseCase.DeleteTransactionID(transactionID)
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
