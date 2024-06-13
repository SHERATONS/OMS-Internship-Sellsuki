package Handlers

import (
	"encoding/json"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
	"github.com/gofiber/fiber/v2"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type ProductHandler struct {
	UseCases UseCases.IProductCase
}

func (s *ProductHandler) GetProductById(c *fiber.Ctx) error {
	productId := c.Params("id")
	if productId == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Product Id is Required")
	}
	product, err := s.UseCases.GetProductById(productId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Invalid Product Id")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"product": product})
}

func (s *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var rawData map[string]interface{}
	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Request Body")
	}

	var validationError []string

	if pId, ok := rawData["PID"]; ok {
		CheckIdString := pId.(string)
		if reflect.TypeOf(pId).Kind() != reflect.String {
			validationError = append(validationError, "Product ID Must Be a String")
		} else if CheckIdInt, _ := strconv.Atoi(CheckIdString); CheckIdInt <= 0 {
			validationError = append(validationError, "Product ID Must Greater than 0")
		}
	} else {
		validationError = append(validationError, "Product ID is Required and Must Be a String")
	}

	if pName, ok := rawData["PName"]; ok {
		if reflect.TypeOf(pName).Kind() != reflect.String {
			validationError = append(validationError, "Product Name Must Be String")
		}
	} else {
		validationError = append(validationError, "Product Name is Required and Must Be String")
	}

	if pPrice, ok := rawData["PPrice"]; ok {
		CheckPriceFloat := pPrice.(float64)
		if reflect.TypeOf(pPrice).Kind() != reflect.Float64 {
			validationError = append(validationError, "Product Price Must Be Float")
		} else if CheckPriceFloat <= 0 {
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
		return c.Status(fiber.StatusBadRequest).SendString(strings.Join(validationError, ", "))
	}

	var createProduct Entities.Product
	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error Processing Request Data")
	}
	if err := json.Unmarshal(data, &createProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error Processing Request Data")
	}

	createProduct.PCreated = time.Now()
	createProduct.PUpdated = time.Now()

	product, err := s.UseCases.CreateProduct(createProduct)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Product ID Already Exists")
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"product": product})
}

func (s *ProductHandler) UpdateProductById(c *fiber.Ctx) error {
	var rawData map[string]interface{}
	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Request Body")
	}

	var validationError []string

	if pId, ok := rawData["PID"]; ok {
		CheckIdString := pId.(string)
		if reflect.TypeOf(pId).Kind() != reflect.String {
			validationError = append(validationError, "Product ID Must Be String")
		} else if CheckIdInt, _ := strconv.Atoi(CheckIdString); CheckIdInt <= 0 {
			validationError = append(validationError, "Product ID Must Be Greater than 0")
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
		CheckPriceFloat := pPrice.(float64)
		if reflect.TypeOf(pPrice).Kind() != reflect.Float64 {
			validationError = append(validationError, "Product Price Must Be Float")
		} else if CheckPriceFloat <= 0 {
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
		return c.Status(fiber.StatusBadRequest).SendString(strings.Join(validationError, ", "))
	}
	var updateProduct Entities.Product
	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error Processing Request Data")
	}
	if err := json.Unmarshal(data, &updateProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error Processing Request Data")
	}

	var productId = c.Params("id")
	product, err := s.UseCases.GetProductById(productId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Product Id Not Found")
	}

	product.PID = updateProduct.PID
	product.PName = updateProduct.PName
	product.PPrice = updateProduct.PPrice
	product.PDesc = updateProduct.PDesc
	product.PUpdated = time.Now()

	updateProduct, err = s.UseCases.UpdateProduct(product, productId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Product Already Exists")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"product": updateProduct})
}

func (s *ProductHandler) DeleteProductById(c *fiber.Ctx) error {
	productId := c.Params("id")
	product, err := s.UseCases.GetProductById(productId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Invalid Product Id")
	}
	err = s.UseCases.DeleteProductById(productId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to Delete Product")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Product successfully deleted",
		"productId": product.PID,
	})
}

func (s *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := s.UseCases.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Something Went Wrong")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"products": products})
}

func NewProductHandler(useCases UseCases.IProductCase) ProductHandlerI {
	return &ProductHandler{UseCases: useCases}
}
