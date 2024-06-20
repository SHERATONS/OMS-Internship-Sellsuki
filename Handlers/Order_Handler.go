package Handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"reflect"
	"strconv"
	"strings"
	"time"

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

	if TranID, ok := rawData["OTranID"].(string); ok {
		if reflect.TypeOf(TranID).Kind() != reflect.String {
			validationError = append(validationError, "Transaction ID Must Be a String")
		}
	} else {
		validationError = append(validationError, "Transaction ID is Required and Must Be a string")
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	TempOrder, err := o.UseCaseTransactionID.GetOrderByTransactionID(rawData["OTranID"].(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Transaction ID Not Found"})
	}

	productList := strings.Split(TempOrder.TProductList, ", ")
	for _, product := range productList {
		parts := strings.Split(product, ":")
		PID := strings.TrimSpace(parts[0])
		PQuantity, _ := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)

		if _, err := o.UseCaseStock.GetStockByID(PID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Stock ID Not Found or Did not Create Stock"})
		}

		StockQuantity, _ := o.UseCaseStock.GetStockByID(PID)
		NewQuantity := StockQuantity.SQuantity - PQuantity

		TempStock := Entities.Stock{
			SID:       PID,
			SQuantity: NewQuantity,
			SUpdated:  time.Now(),
		}

		_, err := o.UseCaseStock.UpdateStock(TempStock, PID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot Create Order, Because Stock is Not Enough"})
		}
	}

	var createOrder Entities.Order

	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}
	if err := json.Unmarshal(data, &createOrder); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}

	createOrder.OID = uuid.New()
	createOrder.OTranID = TempOrder.TID
	createOrder.OPaid = false
	createOrder.ODestination = TempOrder.TDestination
	createOrder.OPrice = TempOrder.TPrice
	createOrder.OStatus = "New"
	createOrder.OCreated = time.Now()

	_, err = o.UseCase.CreateOrder(createOrder)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to Create Order"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Created Order Succeed"})
}

func (o *OrderHandler) ChangeOrderStatus(c *fiber.Ctx) error {
	var rawData map[string]interface{}

	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	if OStatus, ok := rawData["OStatus"].(string); ok {
		if reflect.TypeOf(OStatus).Kind() != reflect.String {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Order Status Must Be String"})
		}
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Order Status is Required and Must Be String"})
	}

	OrderStatus := rawData["OStatus"].(string)

	var Order Entities.Order

	OrderID := c.Params("oid")

	switch OrderStatus {
	case "Paid":
		TempOrder, err := o.UseCase.GetOrderById(OrderID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Order ID Not Found"})
		}
		if TempOrder.OStatus == "New" {
			Order.OStatus = "Paid"
			Order.OPaid = true
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Order Status"})
		}

		NewStatus, err := o.UseCase.ChangeOrderStatus(Order, OrderID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot Change Order Status"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"order": NewStatus})

	case "Processing":
		TempOrder, err := o.UseCase.GetOrderById(OrderID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Order ID Not Found"})
		}

		if TempOrder.OStatus == "Paid" {
			if TempOrder.ODestination != "Branch" {
				Order.OStatus = "Processing"
				Order.OPaid = true
			} else {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Please Come Pick Up your Product at the Branch"})
			}
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Order Status"})
		}

		NewStatus, err := o.UseCase.ChangeOrderStatus(Order, OrderID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot Change Order Status"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"order": NewStatus})

	case "Done":
		TempOrder, err := o.UseCase.GetOrderById(OrderID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Order ID Not Found"})
		}
		if TempOrder.OStatus == "Processing" {
			Order.OStatus = "Done"
			Order.OPaid = true
		} else if TempOrder.OStatus == "Paid" && TempOrder.ODestination == "Branch" {
			Order.OStatus = "Done"
			Order.OPaid = true
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Order Status"})
		}

		NewStatus, err := o.UseCase.ChangeOrderStatus(Order, OrderID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot Change Order Status"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"order": NewStatus})

	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Order Status"})
	}
}

func NewOrderHandler(useCase UseCases.IOrderUseCase, useCaseStock UseCases.IStockUseCase, useCaseTransactionID UseCases.ITransactionIDUseCase) IOrderHandler {
	return &OrderHandler{
		UseCase:              useCase,
		UseCaseStock:         useCaseStock,
		UseCaseTransactionID: useCaseTransactionID,
	}
}
