package Product

import (
	"encoding/json"
	Product2 "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Product"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	UseCase UseCases.IProductUseCase
}

func (s *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	ctx, span := tracer.Start(c.UserContext(), "GetProductByID_Handler")
	defer span.End()

	productID := c.Params("id")

	product, err := s.UseCase.GetProductById(ctx, productID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Product ID Not Found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"product": product})
}

func (s *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	ctx, span := tracer.Start(c.UserContext(), "CreateProduct_Handler")
	defer span.End()

	var rawData map[string]interface{}

	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	var validationError []string

	var tempProduct Product2.Product

	if err := tempProduct.ValidateProductID(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := tempProduct.ValidateProductName(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := tempProduct.ValidateProductPrice(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := tempProduct.ValidateProductDescription(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	var createProduct Product2.Product

	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}
	if err := json.Unmarshal(data, &createProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}

	product, err := s.UseCase.CreateProduct(ctx, createProduct)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Product ID Already Exists"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"product": product})
}

func (s *ProductHandler) UpdateProductById(c *fiber.Ctx) error {
	ctx, span := tracer.Start(c.UserContext(), "UpdateProductById_Handler")
	defer span.End()

	var rawData map[string]interface{}

	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	var validationError []string

	var tempProduct Product2.Product

	if err := tempProduct.ValidateProductID(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := tempProduct.ValidateProductName(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := tempProduct.ValidateProductPrice(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := tempProduct.ValidateProductDescription(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	var updateProduct Product2.Product

	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}
	if err := json.Unmarshal(data, &updateProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}

	var productId = c.Params("id")

	updateProduct, err = s.UseCase.UpdateProduct(ctx, updateProduct, productId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"product": updateProduct})
}

func (s *ProductHandler) DeleteProductById(c *fiber.Ctx) error {
	ctx, span := tracer.Start(c.UserContext(), "DeleteProductById_Handler")
	defer span.End()

	productID := c.Params("id")

	err := s.UseCase.DeleteProductById(ctx, productID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Product Successfully Deleted and Stock Successfully Deleted",
		"productID": productID,
	})
}

func (s *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	ctx, span := tracer.Start(c.UserContext(), "GetAllProducts_Handler")
	defer span.End()

	products, err := s.UseCase.GetAllProducts(ctx)
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
