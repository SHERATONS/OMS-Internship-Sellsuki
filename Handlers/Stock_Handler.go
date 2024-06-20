package Handlers

import (
	"encoding/json"
	"reflect"
	"strconv"
	"time"

	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
	"github.com/gofiber/fiber/v2"
)

type StockHandler struct {
	UseCase        UseCases.IStockUseCase
	UseCaseProduct UseCases.IProductUseCase
}

func (s *StockHandler) DeleteStock(c *fiber.Ctx) error {
	var stockId = c.Params("id")

	stock, err := s.UseCase.GetStockByID(stockId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Stock ID Not Found"})
	}

	err = s.UseCase.DeleteStock(stockId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to Delete Stock"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Stock successfully deleted",
		"productId": stock.SID,
	})
}

func (s *StockHandler) UpdateStock(c *fiber.Ctx) error {
	var rawData map[string]interface{}
	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	var validationError []string

	if sId, ok := rawData["SID"]; ok {
		if reflect.TypeOf(sId).Kind() != reflect.String {
			validationError = append(validationError, "Stock ID Must Be a String")
		} else {
			CheckIdString := sId.(string)
			if CheckIdInt, err := strconv.Atoi(CheckIdString); err != nil {
				validationError = append(validationError, "Stock ID Must a Number")
			} else if CheckIdInt <= 0 {
				validationError = append(validationError, "Stock ID Must Greater than 0")
			} else {
				_, err := s.UseCaseProduct.GetProductById(CheckIdString)
				if err != nil {
					validationError = append(validationError, "Stock ID Did not Exists in Product ID")
				}
			}
		}
	} else {
		validationError = append(validationError, "Stock ID is Required and Must Be a String")
	}

	if sQuantity, ok := rawData["SQuantity"]; ok {
		CheckQuantityInt := sQuantity.(float64)
		if reflect.TypeOf(sQuantity).Kind() != reflect.Float64 {
			validationError = append(validationError, "Stock Quantity Must Be a Integer")
		} else if CheckQuantityInt < 0 {
			validationError = append(validationError, "Stock Quantity Must Be Greater than 0")
		}
	} else {
		validationError = append(validationError, "Stock Quantity is Required and Must Be a Integer")
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	var updateStock Entities.Stock
	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}
	if err := json.Unmarshal(data, &updateStock); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}

	var stockId = c.Params("id")

	stock, err := s.UseCase.GetStockByID(stockId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Stock ID Not Found"})
	}

	stock.SID = updateStock.SID
	stock.SQuantity = updateStock.SQuantity
	stock.SUpdated = time.Now()

	updateStock, err = s.UseCase.UpdateStock(stock, stockId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to Updated Stock"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"stock": updateStock})
}

func (s *StockHandler) CreateStock(c *fiber.Ctx) error {
	var rawData map[string]interface{}

	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	var validationError []string

	if sId, ok := rawData["SID"]; ok {
		if reflect.TypeOf(sId).Kind() != reflect.String {
			validationError = append(validationError, "Stock ID Must Be a String")
		} else {
			CheckIdString := sId.(string)
			if CheckIdInt, err := strconv.Atoi(CheckIdString); err != nil {
				validationError = append(validationError, "Stock ID Must a Number")
			} else if CheckIdInt <= 0 {
				validationError = append(validationError, "Stock ID Must Greater than 0")
			} else {
				_, err := s.UseCaseProduct.GetProductById(CheckIdString)
				if err != nil {
					validationError = append(validationError, "Stock ID Did not Exists in Product ID")
				}
			}
		}
	} else {
		validationError = append(validationError, "Stock ID is Required and Must Be a String")
	}

	if sQuantity, ok := rawData["SQuantity"]; ok {
		CheckQuantityInt := sQuantity.(float64)
		if reflect.TypeOf(sQuantity).Kind() != reflect.Float64 {
			validationError = append(validationError, "Stock Quantity Must Be a Integer")
		} else if CheckQuantityInt <= 0 {
			validationError = append(validationError, "Stock Quantity Must Be Greater than 0")
		}
	} else {
		validationError = append(validationError, "Stock Quantity is Required and Must Be a Integer")
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	var createStock Entities.Stock
	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}
	err = json.Unmarshal(data, &createStock)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}

	createStock.SUpdated = time.Now()

	stock, err := s.UseCase.CreateStock(createStock)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to Created Stock"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"Stock": stock})
}

func (s *StockHandler) GetStockByID(c *fiber.Ctx) error {
	stockID := c.Params("id")

	stock, err := s.UseCase.GetStockByID(stockID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Stock ID Not Found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Stock": stock})
}

func (s *StockHandler) GetAllStock(c *fiber.Ctx) error {
	stocks, err := s.UseCase.GetAllStocks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something Went Wrong"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"stocks": stocks})
}

func NewStockHandler(useCase UseCases.IStockUseCase, useCaseProduct UseCases.IProductUseCase) IStockHandler {
	return &StockHandler{
		UseCase:        useCase,
		UseCaseProduct: useCaseProduct,
	}
}
