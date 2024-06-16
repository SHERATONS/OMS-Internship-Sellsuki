package Handlers

import (
	"encoding/json"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
	"github.com/gofiber/fiber/v2"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type OrderCalculateHandler struct {
	UseCases        UseCases.IOrderCalculateCase
	UseCasesProduct UseCases.IProductCase
	UseCasesAddress UseCases.IAddressCase
}

func (o *OrderCalculateHandler) GetAllOrder(c *fiber.Ctx) error {
	transactionID, err := o.UseCases.GetAllOrders()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": "Something went wrong",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"transaction_id": transactionID,
	})

}

func (o *OrderCalculateHandler) GetOrderByTransactionID(c *fiber.Ctx) error {
	transactionID := c.Params("tid")
	if transactionID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Transaction Id is Required",
		})
	}
	order, err := o.UseCases.GetOrderByTransactionID(transactionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid Transaction Id",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"TransactionID": order})
}

func (o *OrderCalculateHandler) CreateTransactionID(c *fiber.Ctx) error {
	var rawData map[string]interface{}
	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Request Body")
	}

	var validationError []string

	if Destination, ok := rawData["ODestination"].(string); ok {
		if reflect.TypeOf(Destination).Kind() != reflect.String {
			validationError = append(validationError, "Destination Must Be a String")
		}
	} else {
		validationError = append(validationError, "Destination is Required and Must Be a string")
	}

	if Product, ok := rawData["OProduct"].(string); ok {
		if reflect.TypeOf(Product).Kind() != reflect.String {
			validationError = append(validationError, "Product Must Be a String")
		}
	} else {
		validationError = append(validationError, "Product is Required and Must Be a string")
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validationError,
		})
	}

	var totalPrice float64

	productList := strings.Split(rawData["OProduct"].(string), ", ")
	for _, product := range productList {
		parts := strings.Split(product, ":")
		if len(parts) == 2 {
			PID := strings.TrimSpace(parts[0])
			PQuantity, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Invalid Quantity",
				})
			}
			if PQuantity <= 0 {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Quantity Must Greater than 0",
				})
			}

			if temp, err := o.UseCasesProduct.GetProductById(PID); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Invalid Product Id",
				})
			} else {
				totalPrice += temp.PPrice * float64(PQuantity)
			}
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid Product Format",
			})
		}
	}

	Destination := rawData["ODestination"].(string)

	NewDestination, err := url.QueryUnescape(Destination)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Destination Parameter",
		})
	}

	address, err := o.UseCasesAddress.GetAddressByCity(NewDestination)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Destination Not Found",
		})
	}

	totalPrice += address.APrice

	var createTransactionID Entities.OrderCalculate
	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error Processing Data",
		})
	}
	if err := json.Unmarshal(data, &createTransactionID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error Processing Data",
		})
	}

	createTransactionID.OTotalPrice = totalPrice
	createTransactionID.OTranID = createTransactionID.GenerateTransactionID(totalPrice)

	transactionID, err := o.UseCases.CreateTransactionID(createTransactionID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error Processing Data",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"TransactionID": transactionID})
}

func (o *OrderCalculateHandler) DeleteTransactionID(c *fiber.Ctx) error {
	transactionID := c.Params("tid")
	order, err := o.UseCases.GetOrderByTransactionID(transactionID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Transaction Id",
		})
	}

	err = o.UseCases.DeleteTransactionID(transactionID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to Delete Transaction Id",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":        "Transaction Id Deleted Successfully",
		"transaction_id": order,
	})
}

func NewOrderCalculateHandler(useCases UseCases.IOrderCalculateCase, useCasesProduct UseCases.IProductCase, useCasesAddress UseCases.IAddressCase) OrderCalculateHandlerI {
	return &OrderCalculateHandler{
		UseCases:        useCases,
		UseCasesProduct: useCasesProduct,
		UseCasesAddress: useCasesAddress,
	}
}
