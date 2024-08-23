package UseCases

import (
	"context"
	"errors"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Address"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Product"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/TransactionID"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/PKG/UseCases/MockRepository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func setupTransactionIDUseCase() (*MockRepository.MockTransactionIDRepo, *MockRepository.MockProductRepo, *MockRepository.MockAddressRepo, ITransactionIDUseCase) {
	mockTransactionIDRepo := new(MockRepository.MockTransactionIDRepo)
	mockProductRepo := new(MockRepository.MockProductRepo)
	mockAddressRepo := new(MockRepository.MockAddressRepo)
	useCase := NewTransactionIDUseCase(mockTransactionIDRepo, mockProductRepo, mockAddressRepo)

	return mockTransactionIDRepo, mockProductRepo, mockAddressRepo, useCase
}

func TestTransactionIDUseCase(t *testing.T) {
	ctx := context.TODO()

	t.Run("Test GetAllTransactionIDs Success", func(t *testing.T) {
		mockTransactionIDRepo, _, _, useCase := setupTransactionIDUseCase()

		sampleTransactionIDs := []TransactionID.TransactionID{
			{TID: "1", TPrice: 100.00, TDestination: "Sample Destination", TProductList: "1:1"},
			{TID: "2", TPrice: 100.00, TDestination: "Sample Destination", TProductList: "1:1"},
		}

		mockTransactionIDRepo.On("GetAllTransactionIDs", mock.Anything).Return(sampleTransactionIDs, nil)

		result, err := useCase.GetAllTransactionIDs(ctx)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, sampleTransactionIDs, result)
		mockTransactionIDRepo.AssertExpectations(t)
	})

	t.Run("Test GetAllTransactionIDs Failure", func(t *testing.T) {
		mockTransactionIDRepo, _, _, useCase := setupTransactionIDUseCase()

		mockTransactionIDRepo.On("GetAllTransactionIDs", mock.Anything).Return([]TransactionID.TransactionID{}, assert.AnError)

		result, err := useCase.GetAllTransactionIDs(ctx)

		assert.Error(t, err)
		assert.Equal(t, []TransactionID.TransactionID{}, result)
		mockTransactionIDRepo.AssertExpectations(t)
	})

	t.Run("Test GetOrderByTransactionID Success", func(t *testing.T) {
		mockTransactionIDRepo, _, _, useCase := setupTransactionIDUseCase()

		sampleTransactionIDs := TransactionID.TransactionID{
			TID:          "1",
			TPrice:       100.00,
			TDestination: "Sample Destination",
			TProductList: "Sample Product List",
		}

		mockTransactionIDRepo.On("GetOrderByTransactionID", mock.Anything, "1").Return(sampleTransactionIDs, nil)

		result, err := useCase.GetOrderByTransactionID(ctx, "1")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, sampleTransactionIDs, result)
		mockTransactionIDRepo.AssertExpectations(t)
	})

	t.Run("Test GetOrderByTransactionID Failure", func(t *testing.T) {
		mockTransactionIDRepo, _, _, useCase := setupTransactionIDUseCase()

		mockTransactionIDRepo.On("GetOrderByTransactionID", mock.Anything, "1").Return(TransactionID.TransactionID{}, errors.New("transaction ID not found"))

		result, err := useCase.GetOrderByTransactionID(ctx, "1")

		assert.Error(t, err)
		assert.Equal(t, "transaction ID not found", err.Error())
		assert.Equal(t, TransactionID.TransactionID{}, result)
		mockTransactionIDRepo.AssertExpectations(t)
	})

	t.Run("Test CreateTransactionID Success", func(t *testing.T) {
		mockTransactionIDRepo, mockProductRepo, mockAddressRepo, useCase := setupTransactionIDUseCase()

		sampleTransactionIDInfo := TransactionID.TransactionID{
			TDestination: "Branch",
			TProductList: "1:1, 2:1",
		}

		sampleProduct1 := Product.Product{
			PID:    "1",
			PName:  "Sample Product 1",
			PPrice: 100,
			PDesc:  "Sample Description",
		}

		sampleProduct2 := Product.Product{
			PID:    "2",
			PName:  "Sample Product 2",
			PPrice: 100,
			PDesc:  "Sample Description",
		}

		sampleAddress := Address.Address{
			City:    "Branch",
			Country: "Thailand",
			APrice:  0,
		}

		expectedTransaction := sampleTransactionIDInfo
		expectedTransaction.TPrice = 200.00
		expectedTransaction.TID = expectedTransaction.GenerateTransactionID(expectedTransaction.TPrice)

		mockProductRepo.On("GetProductByID", mock.Anything, "1").Return(sampleProduct1, nil)
		mockProductRepo.On("GetProductByID", mock.Anything, "2").Return(sampleProduct2, nil)
		mockAddressRepo.On("GetAddressByCity", mock.Anything, "Branch").Return(sampleAddress, nil)
		mockTransactionIDRepo.On("CreateTransactionID", mock.Anything, expectedTransaction).Return(expectedTransaction, nil)

		result, err := useCase.CreateTransactionID(ctx, sampleTransactionIDInfo)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedTransaction, result)
		mockProductRepo.AssertExpectations(t)
		mockAddressRepo.AssertExpectations(t)
		mockTransactionIDRepo.AssertExpectations(t)
	})

	t.Run("Test CreateTransactionID Failure", func(t *testing.T) {

	})

	t.Run("Test CreateTransactionID Failure Invalid Product List Format", func(t *testing.T) {
		_, _, _, useCase := setupTransactionIDUseCase()

		sampleTransactionInfo := TransactionID.TransactionID{
			TDestination: "Bangkok",
			TProductList: "1:2:3",
		}

		_, err := useCase.CreateTransactionID(ctx, sampleTransactionInfo)

		assert.Error(t, err)
		assert.Equal(t, "invalid Product Format, Should be Like This 'ProductID:Quantity'", err.Error())
	})

	t.Run("Test CreateTransactionID Failure Invalid Quantity", func(t *testing.T) {
		_, _, _, useCase := setupTransactionIDUseCase()

		sampleTransactionInfo := TransactionID.TransactionID{
			TDestination: "Bangkok",
			TProductList: "2:Hello",
		}

		_, err := useCase.CreateTransactionID(ctx, sampleTransactionInfo)

		assert.Error(t, err)
		assert.Equal(t, "invalid Quantity", err.Error())
	})

	t.Run("Test DeleteTransactionID Success", func(t *testing.T) {
		mockTransactionIDRepo, _, _, useCase := setupTransactionIDUseCase()

		mockTransactionIDRepo.On("GetOrderByTransactionID", mock.Anything, "1").Return(TransactionID.TransactionID{}, nil)
		mockTransactionIDRepo.On("DeleteTransactionID", mock.Anything, "1").Return(nil)

		err := useCase.DeleteTransactionID(ctx, "1")

		assert.NoError(t, err)
		mockTransactionIDRepo.AssertExpectations(t)
	})

	t.Run("Test DeleteTransactionID Failure", func(t *testing.T) {
		mockTransactionIDRepo, _, _, useCase := setupTransactionIDUseCase()

		mockTransactionIDRepo.On("GetOrderByTransactionID", mock.Anything, "1").Return(TransactionID.TransactionID{}, nil)
		mockTransactionIDRepo.On("DeleteTransactionID", mock.Anything, "1").Return(errors.New("failed to delete transaction ID"))

		err := useCase.DeleteTransactionID(ctx, "1")

		assert.Error(t, err)
		mockTransactionIDRepo.AssertExpectations(t)
	})
}
