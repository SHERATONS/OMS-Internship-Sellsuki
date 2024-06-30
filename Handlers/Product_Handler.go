package Handlers

import (
	"encoding/json"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	UseCase UseCases.IProductUseCase
}

func (s *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	productID := c.Params("id")

	product, err := s.UseCase.GetProductById(productID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Product ID Not Found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"product": product})
}

func (s *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var rawData map[string]interface{}

	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	var validationError []string

	if err := Entities.ValidateProductID(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := Entities.ValidateProductName(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := Entities.ValidateProductPrice(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := Entities.ValidateProductDescription(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	var createProduct Entities.Product

	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}
	if err := json.Unmarshal(data, &createProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}

	product, err := s.UseCase.CreateProduct(createProduct)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Product ID Already Exists"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"product": product})
}

func (s *ProductHandler) UpdateProductById(c *fiber.Ctx) error {
	var rawData map[string]interface{}

	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	var validationError []string

	if err := Entities.ValidateProductID(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := Entities.ValidateProductName(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := Entities.ValidateProductPrice(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := Entities.ValidateProductDescription(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	var updateProduct Entities.Product

	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}
	if err := json.Unmarshal(data, &updateProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}

	var productId = c.Params("id")

	updateProduct, err = s.UseCase.UpdateProduct(updateProduct, productId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"product": updateProduct})
}

func (s *ProductHandler) DeleteProductById(c *fiber.Ctx) error {
	productID := c.Params("id")

	err := s.UseCase.DeleteProductById(productID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Product Successfully Deleted and Stock Successfully Deleted",
		"productID": productID,
	})
}

func (s *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := s.UseCase.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something Went Wrong"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"products": products})
}

func NewProductHandler(useCase UseCases.IProductUseCase) IProductHandler {
	return &ProductHandler{
		UseCase: useCase,
	}
}
