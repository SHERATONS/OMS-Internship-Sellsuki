package Address

import (
	"encoding/json"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases/Address"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

type AddressHandler struct {
	UseCase Address.IAddressUseCase
}

func (a *AddressHandler) GetAddressByCity(c *fiber.Ctx) error {
	city := c.Params("city")

	NewCity, err := url.QueryUnescape(city)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid City Parameter"})
	}

	address, err := a.UseCase.GetAddressByCity(NewCity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"address": address})
}

func (a *AddressHandler) CreateAddress(c *fiber.Ctx) error {
	var rawData map[string]interface{}
	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	var validationError []string

	if err := Entities.ValidateCity(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := Entities.ValidateCountry(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := Entities.ValidateAPrice(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	var createAddress Entities.Address

	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}
	if err := json.Unmarshal(data, &createAddress); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}

	address, err := a.UseCase.CreateAddress(createAddress)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"address": address})
}

func (a *AddressHandler) UpdateAddress(c *fiber.Ctx) error {
	var rawData map[string]interface{}

	if err := c.BodyParser(&rawData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	var validationError []string

	if err := Entities.ValidateCity(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := Entities.ValidateCountry(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if err := Entities.ValidateAPrice(rawData); err != nil {
		validationError = append(validationError, err.Error())
	}

	if len(validationError) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	var updateAddress Entities.Address

	data, err := json.Marshal(rawData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}
	if err := json.Unmarshal(data, &updateAddress); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error Processing Request Data"})
	}

	city := c.Params("city")

	NewCity, err := url.QueryUnescape(city)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid City Parameter"})
	}

	updateAddress, err = a.UseCase.UpdateAddress(updateAddress, NewCity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"address": updateAddress})
}

func (a *AddressHandler) DeleteAddress(c *fiber.Ctx) error {
	city := c.Params("city")

	NewCity, err := url.QueryUnescape(city)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid City Parameter"})
	}

	err = a.UseCase.DeleteAddress(NewCity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Address successfully deleted",
		"city":    NewCity,
	})
}

func NewAddressHandler(useCase Address.IAddressUseCase) IAddressHandler {
	return &AddressHandler{UseCase: useCase}
}
