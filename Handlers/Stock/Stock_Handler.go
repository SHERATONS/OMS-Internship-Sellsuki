package Stock

import (
	"encoding/json"
	Stock2 "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Stock"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
	"github.com/gofiber/fiber/v2"
)

type StockHandler struct {
	UseCase UseCases.IStockUseCase
}

func (s *StockHandler) DeleteStock(c *fiber.Ctx) error {
	ctx, span := tracer.Start(c.UserContext(), "DeleteStock_Handler")
	defer span.End()

	var stockId = c.Params("id")

	err := s.UseCase.DeleteStock(ctx, stockId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Stock successfully deleted",
		"productId": stockId,
	})
}

func (s *StockHandler) UpdateStock(c *fiber.Ctx) error {
	ctx, span := tracer.Start(c.UserContext(), "UpdateStock_Handler")
	defer span.End()

	var rawData map[string]interface{}
	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	var validationError []string

	var tempStock Stock2.Stock

	if err := tempStock.ValidateStockID(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := tempStock.ValidateStockQuantity(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	var updateStock Stock2.Stock

	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}
	if err := json.Unmarshal(data, &updateStock); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}

	var stockId = c.Params("id")

	updateStock, err = s.UseCase.UpdateStock(ctx, updateStock, stockId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"stock": updateStock})
}

func (s *StockHandler) CreateStock(c *fiber.Ctx) error {
	ctx, span := tracer.Start(c.UserContext(), "CreateStock_Handler")
	defer span.End()

	var rawData map[string]interface{}

	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	var validationError []string

	var tempStock Stock2.Stock

	if err := tempStock.ValidateStockID(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := tempStock.ValidateStockQuantity(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	var createStock Stock2.Stock

	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}
	err = json.Unmarshal(data, &createStock)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}

	stock, err := s.UseCase.CreateStock(ctx, createStock)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"Stock": stock})
}

func (s *StockHandler) GetStockByID(c *fiber.Ctx) error {
	ctx, span := tracer.Start(c.UserContext(), "GetStockByID_Handler")
	defer span.End()

	stockID := c.Params("id")

	stock, err := s.UseCase.GetStockByID(ctx, stockID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Stock ID Not Found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Stock": stock})
}

func (s *StockHandler) GetAllStock(c *fiber.Ctx) error {
	ctx, span := tracer.Start(c.UserContext(), "GetAllStock_Handler")
	defer span.End()

	stocks, err := s.UseCase.GetAllStocks(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something Went Wrong"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"stocks": stocks})
}

func NewStockHandler(useCase UseCases.IStockUseCase) IStockHandler {
	return &StockHandler{UseCase: useCase}
}
