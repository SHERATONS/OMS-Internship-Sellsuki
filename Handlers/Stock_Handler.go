package Handlers

import (
	"encoding/json"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
	"github.com/gofiber/fiber/v2"
)

type StockHandler struct {
	UseCase UseCases.IStockUseCase
}

func (s *StockHandler) DeleteStock(c *fiber.Ctx) error {
	var stockId = c.Params("id")

	err := s.UseCase.DeleteStock(stockId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Stock successfully deleted",
		"productId": stockId,
	})
}

func (s *StockHandler) UpdateStock(c *fiber.Ctx) error {
	var rawData map[string]interface{}
	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	var validationError []string

	if err := Entities.ValidateStockID(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := Entities.ValidateStockQuantity(rawData); err != nil {
		validationError = append(validationError, err.Error())
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

	updateStock, err = s.UseCase.UpdateStock(updateStock, stockId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"stock": updateStock})
}

func (s *StockHandler) CreateStock(c *fiber.Ctx) error {
	var rawData map[string]interface{}

	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	var validationError []string

	if err := Entities.ValidateStockID(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := Entities.ValidateStockQuantity(rawData); err != nil {
		validationError = append(validationError, err.Error())
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

	stock, err := s.UseCase.CreateStock(createStock)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
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

func NewStockHandler(useCase UseCases.IStockUseCase) IStockHandler {
	return &StockHandler{UseCase: useCase}
}
