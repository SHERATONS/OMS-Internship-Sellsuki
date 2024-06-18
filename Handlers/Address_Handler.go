package Handlers

import (
	"encoding/json"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
	"github.com/gofiber/fiber/v2"
	"net/url"
	"reflect"
	"time"
)

type AddressHandler struct {
	UseCases UseCases.IAddressCase
}

func (a *AddressHandler) GetAddressByCity(c *fiber.Ctx) error {
	city := c.Params("city")
	if city == "" {
		return c.Status(fiber.StatusBadRequest).SendString("City is Required")
	}

	NewCity, err := url.QueryUnescape(city)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid City Parameter")
	}

	address, err := a.UseCases.GetAddressByCity(NewCity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Invalid City")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"address": address})
}

func (a *AddressHandler) CreateAddress(c *fiber.Ctx) error {
	var rawData map[string]interface{}
	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Request Body")
	}

	var validationError []string

	if city, ok := rawData["City"]; ok {
		if reflect.TypeOf(city).Kind() != reflect.String {
			validationError = append(validationError, "City Must Be a String")
		}
	} else {
		validationError = append(validationError, "City is Required and Must Be a String")
	}

	if country, ok := rawData["Country"]; ok {
		if reflect.TypeOf(country).Kind() != reflect.String {
			validationError = append(validationError, "Country Must Be a String")
		}
	} else {
		validationError = append(validationError, "Country is Required and Must Be a String")
	}

	if aPrice, ok := rawData["APrice"]; ok {
		if reflect.TypeOf(aPrice).Kind() != reflect.Float64 {
			validationError = append(validationError, "Price Must Be a Float")
		} else if CheckPriceFloat := aPrice.(float64); CheckPriceFloat <= 0 {
			validationError = append(validationError, "Price Must Greater than 0")
		}
	} else {
		validationError = append(validationError, "Price is Required and Must Be a Float")
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validationError,
		})
	}

	var createAddress Entities.Address
	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error Processing Request Data")
	}
	if err := json.Unmarshal(data, &createAddress); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error Processing Request Data")
	}

	createAddress.AUpdated = time.Now()

	address, err := a.UseCases.CreateAddress(createAddress)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error Processing Request Data")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"address": address})
}

func (a *AddressHandler) UpdateAddress(c *fiber.Ctx) error {
	var rawData map[string]interface{}
	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Request Body")
	}

	var validationError []string

	if city, ok := rawData["City"]; ok {
		if reflect.TypeOf(city).Kind() != reflect.String {
			validationError = append(validationError, "City Must Be a String")
		}
	} else {
		validationError = append(validationError, "City is Required and Must Be a String")
	}

	if country, ok := rawData["Country"]; ok {
		if reflect.TypeOf(country).Kind() != reflect.String {
			validationError = append(validationError, "Country Must Be a String")
		}
	} else {
		validationError = append(validationError, "Country is Required and Must Be a String")
	}

	if APrice, ok := rawData["APrice"]; ok {
		if reflect.TypeOf(APrice).Kind() != reflect.Float64 {
			validationError = append(validationError, "Price Must Be a Float")
		} else if CheckPriceFloat := APrice.(float64); CheckPriceFloat <= 0 {
			validationError = append(validationError, "Price Must Greater than 0")
		}
	} else {
		validationError = append(validationError, "Price is Required and Must Be a Float")
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validationError,
		})
	}

	var updateAddress Entities.Address
	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error Processing Request Data")
	}
	if err := json.Unmarshal(data, &updateAddress); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error Processing Request Data")
	}

	city := c.Params("city")
	if city == "" {
		return c.Status(fiber.StatusBadRequest).SendString("City is Required")
	}

	NewCity, err := url.QueryUnescape(city)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid City Parameter")
	}

	address, err := a.UseCases.GetAddressByCity(NewCity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Invalid City")
	}

	address.City = updateAddress.City
	address.Country = updateAddress.Country
	address.APrice = updateAddress.APrice
	address.AUpdated = time.Now()

	updateAddress, err = a.UseCases.UpdateAddress(address, NewCity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Address Already Exists")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"address": updateAddress})
}

func (a *AddressHandler) DeleteAddress(c *fiber.Ctx) error {
	city := c.Params("city")
	if city == "" {
		return c.Status(fiber.StatusBadRequest).SendString("City is Required")
	}

	NewCity, err := url.QueryUnescape(city)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid City Parameter")
	}

	address, err := a.UseCases.GetAddressByCity(NewCity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Invalid City")
	}

	err = a.UseCases.DeleteAddress(NewCity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to Delete Address")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Address successfully deleted",
		"city":    address.City,
	})
}

func NewAddressHandler(useCases UseCases.IAddressCase) AddressHandlerI {
	return &AddressHandler{UseCases: useCases}
}
