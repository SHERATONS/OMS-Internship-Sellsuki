package Handlers

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
	"github.com/gofiber/fiber/v2"
)

type TransactionIDHandler struct {
	UseCase        UseCases.ITransactionIDUseCase
	UseCaseProduct UseCases.IProductUseCase
	UseCaseAddress UseCases.IAddressUseCase
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

	if Destination, ok := rawData["TDestination"].(string); ok {
		if reflect.TypeOf(Destination).Kind() != reflect.String {
			validationError = append(validationError, "Destination Must Be a String")
		}
	} else {
		validationError = append(validationError, "Destination is Required and Must Be a string")
	}

	if Product, ok := rawData["TProductList"].(string); ok {
		if reflect.TypeOf(Product).Kind() != reflect.String {
			validationError = append(validationError, "Product Must Be a String")
		}
	} else {
		validationError = append(validationError, "Product is Required and Must Be a string")
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	var totalPrice float64
	var tempProductList []string

	productList := strings.Split(rawData["TProductList"].(string), ", ")
	for _, product := range productList {
		parts := strings.Split(product, ":")

		if len(parts) == 2 {
			PID := strings.TrimSpace(parts[0])
			PQuantity, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Quantity"})
			}

			if PQuantity <= 0 {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Quantity Must Greater than 0"})
			}

			for _, id := range tempProductList {
				if id == PID {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Product ID Must Not Duplicated"})
				}
			}

			tempProductList = append(tempProductList, PID)

			if temp, err := o.UseCaseProduct.GetProductById(PID); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Product Id Not Found"})
			} else {
				totalPrice += temp.PPrice * float64(PQuantity)
			}
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Product Format, Should be Like This 'ProductID:Quantity'"})
		}
	}

	Destination := rawData["TDestination"].(string)

	NewDestination, err := url.QueryUnescape(Destination)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Destination Parameter"})
	}

	address, err := o.UseCaseAddress.GetAddressByCity(NewDestination)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Destination Not Found"})
	}

	totalPrice += address.APrice

	var createTransactionID Entities.TransactionID

	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}
	if err := json.Unmarshal(data, &createTransactionID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}

	createTransactionID.TPrice = totalPrice
	createTransactionID.TID = createTransactionID.GenerateTransactionID(totalPrice)

	transactionID, err := o.UseCase.CreateTransactionID(createTransactionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to Create TransactionID"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"TransactionID": transactionID})
}

func (o *TransactionIDHandler) DeleteTransactionID(c *fiber.Ctx) error {
	transactionID := c.Params("tid")

	order, err := o.UseCase.GetOrderByTransactionID(transactionID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Transaction ID Not Found"})
	}

	err = o.UseCase.DeleteTransactionID(transactionID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to Delete Transaction Id"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":        "Transaction Id Deleted Successfully",
		"transaction_id": order,
	})
}

func NewTransactionIDHandler(useCase UseCases.ITransactionIDUseCase, useCaseProduct UseCases.IProductUseCase, useCaseAddress UseCases.IAddressUseCase) ITransactionIDHandler {
	return &TransactionIDHandler{
		UseCase:        useCase,
		UseCaseProduct: useCaseProduct,
		UseCaseAddress: useCaseAddress,
	}
}
