package UseCases

import (
	"context"
	"errors"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Address"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases/MockRepository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func setupAddressUseCase() (*MockRepository.MockAddressRepo, IAddressUseCase) {
	mockAddressRepo := new(MockRepository.MockAddressRepo)
	useCase := NewAddressUseCase(mockAddressRepo)

	return mockAddressRepo, useCase
}

func TestAddressUseCase(t *testing.T) {
	ctx := context.TODO()

	t.Run("Test GetAddressByCity Success", func(t *testing.T) {
		mockAddressRepo, useCase := setupAddressUseCase()

		sampleAddress := Address.Address{
			City:     "Sample City",
			Country:  "Sample Country",
			APrice:   100.0,
			AUpdated: time.Now(),
		}

		mockAddressRepo.On("GetAddressByCity", mock.Anything, "Sample City").Return(sampleAddress, nil)

		result, err := useCase.GetAddressByCity(ctx, "Sample City")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, sampleAddress, result)
		mockAddressRepo.AssertExpectations(t)
	})

	t.Run("Test GetAddressByCity Failure", func(t *testing.T) {
		mockAddressRepo, useCase := setupAddressUseCase()

		mockAddressRepo.On("GetAddressByCity", mock.Anything, "Sample City").Return(Address.Address{}, errors.New("address not found"))

		result, err := useCase.GetAddressByCity(ctx, "Sample City")

		assert.Error(t, err)
		assert.Equal(t, "address not found", err.Error())
		assert.Equal(t, Address.Address{}, result)
		mockAddressRepo.AssertExpectations(t)
	})

	t.Run("Test CreateAddress Success", func(t *testing.T) {
		mockAddressRepo, useCase := setupAddressUseCase()

		sampleAddress := Address.Address{
			City:     "Sample City",
			Country:  "Sample Country",
			APrice:   100.0,
			AUpdated: time.Now(),
		}

		mockAddressRepo.On("CreateAddress", mock.Anything, sampleAddress).Return(sampleAddress, nil)

		result, err := useCase.CreateAddress(ctx, sampleAddress)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, sampleAddress, result)
		mockAddressRepo.AssertExpectations(t)
	})

	t.Run("Test CreateAddress Failure", func(t *testing.T) {
		mockAddressRepo, useCase := setupAddressUseCase()

		sampleAddress := Address.Address{
			City:     "Sample City",
			Country:  "Sample Country",
			APrice:   100.0,
			AUpdated: time.Now(),
		}

		mockAddressRepo.On("CreateAddress", mock.Anything, sampleAddress).Return(Address.Address{}, errors.New("failed to create address"))

		result, err := useCase.CreateAddress(ctx, sampleAddress)

		assert.Error(t, err)
		assert.Equal(t, "failed to create address", err.Error())
		assert.Equal(t, Address.Address{}, result)
		mockAddressRepo.AssertExpectations(t)
	})

	t.Run("Test UpdateAddress Success", func(t *testing.T) {
		mockAddressRepo, useCase := setupAddressUseCase()

		updatedAddress := Address.Address{
			City:     "Update City",
			Country:  "Update Country",
			APrice:   100.0,
			AUpdated: time.Now(),
		}

		mockAddressRepo.On("GetAddressByCity", mock.Anything, "Update City").Return(updatedAddress, nil)
		mockAddressRepo.On("UpdateAddress", mock.Anything, updatedAddress, "Update City").Return(updatedAddress, nil)

		result, err := useCase.UpdateAddress(ctx, updatedAddress, "Update City")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, updatedAddress, result)
		mockAddressRepo.AssertExpectations(t)
	})

	t.Run("Test UpdateAddress Failure", func(t *testing.T) {
		mockAddressRepo, useCase := setupAddressUseCase()

		updatedAddress := Address.Address{
			City:     "Update City",
			Country:  "Update Country",
			APrice:   100.0,
			AUpdated: time.Now(),
		}

		mockAddressRepo.On("GetAddressByCity", mock.Anything, "Update City").Return(updatedAddress, nil)
		mockAddressRepo.On("UpdateAddress", mock.Anything, updatedAddress, "Update City").Return(Address.Address{}, errors.New("failed to update address"))

		result, err := useCase.UpdateAddress(ctx, updatedAddress, "Update City")

		assert.Error(t, err)
		assert.Equal(t, "failed to update address", err.Error())
		assert.Equal(t, Address.Address{}, result)
		mockAddressRepo.AssertExpectations(t)
	})

	t.Run("Test DeleteAddress Success", func(t *testing.T) {
		mockAddressRepo, useCase := setupAddressUseCase()

		mockAddressRepo.On("GetAddressByCity", mock.Anything, "Sample City").Return(Address.Address{}, nil)
		mockAddressRepo.On("DeleteAddress", mock.Anything, "Sample City").Return(nil)

		err := useCase.DeleteAddress(ctx, "Sample City")

		assert.NoError(t, err)
		mockAddressRepo.AssertExpectations(t)
	})

	t.Run("Test DeleteAddress Failure", func(t *testing.T) {
		mockAddressRepo, useCase := setupAddressUseCase()

		mockAddressRepo.On("GetAddressByCity", mock.Anything, "Sample City").Return(Address.Address{}, nil)
		mockAddressRepo.On("DeleteAddress", mock.Anything, "Sample City").Return(errors.New("failed to delete address"))

		err := useCase.DeleteAddress(ctx, "Sample City")

		assert.Error(t, err)
		assert.Equal(t, "failed to delete address", err.Error())
		mockAddressRepo.AssertExpectations(t)
	})

}
