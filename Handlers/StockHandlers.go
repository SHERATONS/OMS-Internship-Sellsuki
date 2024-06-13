package Handlers

import (
	"encoding/json"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
	"github.com/gofiber/fiber/v2"
	"reflect"
	"strings"
	"time"
)

type StockHandler struct {
	UseCases        UseCases.IStockCase
	UseCasesProduct UseCases.IProductCase
}

func (s *StockHandler) CreateStock(c *fiber.Ctx) error {
	var rawData map[string]interface{}
	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Request Body")
	}

	var validationError []string

	if sId, ok := rawData["SID"]; ok {
		if reflect.TypeOf(sId).Kind() != reflect.String {
			validationError = append(validationError, "Stock ID Must Be a String")
		}
	} else {
		validationError = append(validationError, "Stock ID is Required and Must Be a String")
	}

	//if sQuantity, ok := rawData["SQuantity"]; ok {
	//	if reflect.TypeOf(sQuantity).Kind() != reflect.Int8 {
	//		validationError = append(validationError, "Stock Quantity Must Be a Integer")
	//	}
	//} else {
	//	validationError = append(validationError, "Stock Quantity is Required and Must Be a Integer")
	//}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).SendString(strings.Join(validationError, ", "))
	}

	var createStock Entities.Stock
	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error Processing Request Data")
	}
	err = json.Unmarshal(data, &createStock)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error Processing Request Data")
	}

	createStock.SUpdated = time.Now()

	//_, err = s.UseCasesProduct.GetProductById(rawData["SID"].(string))
	//if err == nil {
	//	return c.Status(fiber.StatusBadRequest).SendString("Product Id not Exists")
	//}

	stock, err := s.UseCases.CreateStock(createStock)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Stock Id Already Exists")
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"Stock": stock})

}

func (s *StockHandler) GetStockByID(c *fiber.Ctx) error {
	stockID := c.Params("id")
	if stockID == "" {
		return c.Status(fiber.StatusNotFound).SendString("Stock Id is Required")
	}
	stock, err := s.UseCases.GetStockByID(stockID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Invalid Stock Id")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Stock": stock})
}

func (s *StockHandler) GetAllStock(c *fiber.Ctx) error {
	stocks, err := s.UseCases.GetAllStocks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Something Went Wrong")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"stocks": stocks})
}

func NewStockHandler(useCases UseCases.IStockCase) StockHandlerI {
	return &StockHandler{UseCases: useCases}
}
