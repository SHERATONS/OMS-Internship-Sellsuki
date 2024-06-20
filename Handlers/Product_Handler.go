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

type ProductHandler struct {
	UseCase      UseCases.IProductUseCase
	UseCaseStock UseCases.IStockUseCase
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

	if pId, ok := rawData["PID"]; ok {
		if reflect.TypeOf(pId).Kind() != reflect.String {
			validationError = append(validationError, "Product ID Must Be String")
		} else {
			CheckIdString := pId.(string)
			if CheckIdInt, err := strconv.Atoi(CheckIdString); err != nil {
				validationError = append(validationError, "Product ID Must Be a Number")
			} else if CheckIdInt <= 0 {
				validationError = append(validationError, "Product ID Must Be Greater than 0")
			}
		}
	} else {
		validationError = append(validationError, "Product ID is Required and Must Be String")
	}

	if pName, ok := rawData["PName"]; ok {
		if reflect.TypeOf(pName).Kind() != reflect.String {
			validationError = append(validationError, "Product Name Must Be String")
		}
	} else {
		validationError = append(validationError, "Product Name is Required and Must Be String")
	}

	if pPrice, ok := rawData["PPrice"]; ok {
		if reflect.TypeOf(pPrice).Kind() != reflect.Float64 {
			validationError = append(validationError, "Product Price Must Be Float")
		} else {
			CheckPriceFloat := pPrice.(float64)
			if CheckPriceFloat <= 0 {
				validationError = append(validationError, "Product Price Must Be Greater than 0")
			}
		}
	} else {
		validationError = append(validationError, "Product Price is Required and Must Be Float")
	}

	if pDesc, ok := rawData["PDesc"]; ok {
		if reflect.TypeOf(pDesc).Kind() != reflect.String {
			validationError = append(validationError, "Product Description Must Be String")
		}
	} else {
		validationError = append(validationError, "Product Description is Required and Must Be String")
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

	createProduct.PCreated = time.Now()
	createProduct.PUpdated = time.Now()

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

	if pId, ok := rawData["PID"]; ok {
		if reflect.TypeOf(pId).Kind() != reflect.String {
			validationError = append(validationError, "Product ID Must Be String")
		} else {
			CheckIdString := pId.(string)
			if CheckIdInt, err := strconv.Atoi(CheckIdString); err != nil {
				validationError = append(validationError, "Product ID Must Be a Number")
			} else if CheckIdInt <= 0 {
				validationError = append(validationError, "Product ID Must Be Greater than 0")
			}
		}
	} else {
		validationError = append(validationError, "Product ID is Required and Must Be String")
	}

	if pName, ok := rawData["PName"]; ok {
		if reflect.TypeOf(pName).Kind() != reflect.String {
			validationError = append(validationError, "Product Name Must Be String")
		}
	} else {
		validationError = append(validationError, "Product Name is Required and Must Be String")
	}

	if pPrice, ok := rawData["PPrice"]; ok {
		if reflect.TypeOf(pPrice).Kind() != reflect.Float64 {
			validationError = append(validationError, "Product Price Must Be Float")
		} else if CheckPriceFloat := pPrice.(float64); CheckPriceFloat <= 0 {
			validationError = append(validationError, "Product Price Must Be Greater than 0")
		}
	} else {
		validationError = append(validationError, "Product Price is Required and Must Be Float")
	}

	if pDesc, ok := rawData["PDesc"]; ok {
		if reflect.TypeOf(pDesc).Kind() != reflect.String {
			validationError = append(validationError, "Product Description Must Be String")
		}
	} else {
		validationError = append(validationError, "Product Description is Required and Must Be String")
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

	product, err := s.UseCase.GetProductById(productId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Product ID Not Found"})
	}

	product.PID = updateProduct.PID
	product.PName = updateProduct.PName
	product.PPrice = updateProduct.PPrice
	product.PDesc = updateProduct.PDesc
	product.PUpdated = time.Now()

	updateProduct, err = s.UseCase.UpdateProduct(product, productId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to Updated Product"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"product": updateProduct})
}

func (s *ProductHandler) DeleteProductById(c *fiber.Ctx) error {
	productID := c.Params("id")

	product, err := s.UseCase.GetProductById(productID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Product ID Not Found"})
	}

	err = s.UseCase.DeleteProductById(productID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to Deleted Product"})
	}

	_, err = s.UseCaseStock.GetStockByID(productID)
	if err == nil {
		err = s.UseCaseStock.DeleteStock(productID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to Deleted Stock"})
		} else {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message":   "Product Successfully Deleted and Stock Successfully Deleted",
				"productID": product.PID,
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Product Successfully Deleted",
		"productID": product.PID,
	})
}

func (s *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := s.UseCase.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something Went Wrong"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"products": products})
}

func NewProductHandler(useCase UseCases.IProductUseCase, useCaseStock UseCases.IStockUseCase) IProductHandler {
	return &ProductHandler{
		UseCase:      useCase,
		UseCaseStock: useCaseStock,
	}
}
