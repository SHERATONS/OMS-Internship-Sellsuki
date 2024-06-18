package Handlers

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	UseCases               UseCases.IOrderCase
	UseCasesStock          UseCases.IStockCase
	UseCasesOrderCalculate UseCases.ITransactionIDCase
}

func (o *OrderHandler) GetOrderById(c *fiber.Ctx) error {
	orderId, _ := strconv.ParseInt(c.Params("oid"), 10, 64)
	order, err := o.UseCases.GetOrderById(orderId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Order Id",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"order": order})
}

func (o *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var rawData map[string]interface{}
	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Request Body",
		})
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validationError,
		})
	}

	TempOrder, err := o.UseCasesOrderCalculate.GetOrderByTransactionID(rawData["OTranID"].(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Transaction ID",
		})
	}

	productList := strings.Split(TempOrder.OProduct, ", ")
	for _, product := range productList {
		parts := strings.Split(product, ":")
		PID := strings.TrimSpace(parts[0])
		PQuantity, _ := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)

		if _, err := o.UseCasesStock.GetStockByID(PID); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Stock Did Not Create",
			})
		}

		StockQuantity, _ := o.UseCasesStock.GetStockByID(PID)
		NewQuantity := StockQuantity.SQuantity - PQuantity

		TempStock := Entities.Stock{
			SID:       PID,
			SQuantity: NewQuantity,
			SUpdated:  time.Now(),
		}

		_, err := o.UseCasesStock.UpdateStock(TempStock, PID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot Create Order, Because Stock is Not Enough",
			})
		}
	}

	var createOrder Entities.Order

	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error Processing Request Data",
		})
	}
	if err := json.Unmarshal(data, &createOrder); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error Processing Request Data",
		})
	}

	createOrder.OTranID = TempOrder.OTranID
	createOrder.OPaid = false
	createOrder.ODestination = TempOrder.ODestination
	createOrder.OPrice = TempOrder.OTotalPrice
	createOrder.OStatus = "New"
	createOrder.OCreated = time.Now()

	_, err = o.UseCases.CreateOrder(createOrder)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to Create Order",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Created Order Succeed",
	})
}

func (o *OrderHandler) ChangeOrderStatus(c *fiber.Ctx) error {
	var rawData map[string]interface{}
	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Request Body",
		})
	}

	if OStatus, ok := rawData["OStatus"].(string); ok {
		if reflect.TypeOf(OStatus).Kind() != reflect.String {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Order Status Must Be String",
			})
		}
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Order Status is Required and Must Be String",
		})
	}

	OrderStatus := rawData["OStatus"].(string)
	var Order Entities.Order
	OrderId, _ := strconv.ParseInt(c.Params("oid"), 10, 64)

	switch OrderStatus {
	case "Paid":
		TempOrder, err := o.UseCases.GetOrderById(OrderId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid Order Id",
			})
		}
		if TempOrder.OStatus == "New" {
			Order.OStatus = "Paid"
			Order.OPaid = true
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid Order Status",
			})
		}
		NewStatus, err := o.UseCases.ChangeOrderStatus(Order, OrderId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot Change Order Status",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"order": NewStatus,
		})
	case "Processing":
		TempOrder, err := o.UseCases.GetOrderById(OrderId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid Order Id",
			})
		}
		if TempOrder.OStatus == "Paid" {
			if TempOrder.ODestination != "Branch" {
				Order.OStatus = "Processing"
				Order.OPaid = true
			} else {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Please Come Pick Up your Product at the Branch",
				})
			}
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid Order Status",
			})
		}
		NewStatus, err := o.UseCases.ChangeOrderStatus(Order, OrderId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot Change Order Status",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"order": NewStatus,
		})
	case "Done":
		TempOrder, err := o.UseCases.GetOrderById(OrderId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid Order Id",
			})
		}
		if TempOrder.OStatus == "Processing" {
			Order.OStatus = "Done"
			Order.OPaid = true
		} else if TempOrder.OStatus == "Paid" && TempOrder.ODestination == "Branch" {
			Order.OStatus = "Done"
			Order.OPaid = true
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid Order Status",
			})
		}
		NewStatus, err := o.UseCases.ChangeOrderStatus(Order, OrderId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot Change Order Status",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"order": NewStatus,
		})
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Order Status",
		})
	}
}

func NewOrderHandler(useCases UseCases.IOrderCase, useCasesStock UseCases.IStockCase, useCasesOrderCalculate UseCases.ITransactionIDCase) OrderHandlerI {
	return &OrderHandler{
		UseCases:               useCases,
		UseCasesStock:          useCasesStock,
		UseCasesOrderCalculate: useCasesOrderCalculate,
	}
}
