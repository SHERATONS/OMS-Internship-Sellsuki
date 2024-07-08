package UseCases

import (
	"errors"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Order"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Stock"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/TransactionID"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases/MockRepository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"testing"
	"time"
)

func setupOrderUseCase() (*MockRepository.MockOrderRepo, *MockRepository.MockStockRepo, *MockRepository.MockTransactionIDRepo, IOrderUseCase) {
	mockOrderRepo := new(MockRepository.MockOrderRepo)
	mockStockRepo := new(MockRepository.MockStockRepo)
	mockTransactionIDRepo := new(MockRepository.MockTransactionIDRepo)
	useCase := NewOrderUseCase(mockOrderRepo, mockStockRepo, mockTransactionIDRepo)

	return mockOrderRepo, mockStockRepo, mockTransactionIDRepo, useCase
}

func TestOrderUseCase(t *testing.T) {
	ctx := context.TODO()

	t.Run("Test GetOrderById Success", func(t *testing.T) {
		mockOrderRepo, _, _, useCase := setupOrderUseCase()

		sampleOrder := Order.Order{
			OID:          uuid.New(),
			OTranID:      "sample transactionID",
			OPrice:       100.00,
			ODestination: "sample Destination",
			OStatus:      "New",
			OPaid:        false,
			OCreated:     time.Now(),
		}

		mockOrderRepo.On("GetOrderByID", mock.Anything, sampleOrder.OID.String()).Return(sampleOrder, nil)

		result, err := useCase.GetOrderById(ctx, sampleOrder.OID.String())

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, sampleOrder, result)
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("Test GetOrderById Failure", func(t *testing.T) {
		mockOrderRepo, _, _, useCase := setupOrderUseCase()

		mockOrderRepo.On("GetOrderByID", mock.Anything, "1").Return(Order.Order{}, errors.New("order not found"))

		result, err := useCase.GetOrderById(ctx, "1")

		assert.Error(t, err)
		assert.Equal(t, "order not found", err.Error())
		assert.Equal(t, Order.Order{}, result)
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("Test CreateOrder Success", func(t *testing.T) {
		mockOrderRepo, mockStockRepo, mockTransactionIDRepo, useCase := setupOrderUseCase()

		sampleTransactionID := TransactionID.TransactionID{
			TID:          "1",
			TPrice:       100.00,
			TDestination: "Bangkok",
			TProductList: "1:1, 2:1",
		}

		sampleStock1 := Stock.Stock{
			SID:       "1",
			SQuantity: 100,
		}

		sampleStock2 := Stock.Stock{
			SID:       "2",
			SQuantity: 100,
		}

		createdOrder := Order.Order{
			OID:          uuid.New(),
			OTranID:      "1",
			OPrice:       100.00,
			ODestination: "Bangkok",
			OStatus:      "New",
			OPaid:        false,
			OCreated:     time.Now(),
		}

		mockTransactionIDRepo.On("GetOrderByTransactionID", mock.Anything, "1").Return(sampleTransactionID, nil)
		mockStockRepo.On("GetStockByID", mock.Anything, "1").Return(sampleStock1, nil)
		mockStockRepo.On("GetStockByID", mock.Anything, "2").Return(sampleStock2, nil)
		mockStockRepo.On("UpdateStock", mock.Anything, mock.AnythingOfType("Stock.Stock"), "1").Return(sampleStock1, nil)
		mockStockRepo.On("UpdateStock", mock.Anything, mock.AnythingOfType("Stock.Stock"), "2").Return(sampleStock2, nil)
		mockOrderRepo.On("CreateOrder", mock.Anything, mock.AnythingOfType("Order.Order")).Return(createdOrder, nil)

		result, err := useCase.CreateOrder(ctx, "1")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, createdOrder, result)
		mockTransactionIDRepo.AssertExpectations(t)
		mockStockRepo.AssertExpectations(t)
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("Test CreateOrder Failure", func(t *testing.T) {
		mockOrderRepo, mockStockRepo, mockTransactionIDRepo, useCase := setupOrderUseCase()

		sampleTransactionID := TransactionID.TransactionID{
			TID:          "1",
			TPrice:       100.00,
			TDestination: "Bangkok",
			TProductList: "1:1, 2:1",
		}

		sampleStock1 := Stock.Stock{
			SID:       "1",
			SQuantity: 100,
		}

		sampleStock2 := Stock.Stock{
			SID:       "2",
			SQuantity: 100,
		}

		mockTransactionIDRepo.On("GetOrderByTransactionID", mock.Anything, "1").Return(sampleTransactionID, nil)
		mockStockRepo.On("GetStockByID", mock.Anything, "1").Return(sampleStock1, nil)
		mockStockRepo.On("GetStockByID", mock.Anything, "2").Return(sampleStock2, nil)
		mockStockRepo.On("UpdateStock", mock.Anything, mock.AnythingOfType("Stock.Stock"), "1").Return(sampleStock1, nil)
		mockStockRepo.On("UpdateStock", mock.Anything, mock.AnythingOfType("Stock.Stock"), "2").Return(sampleStock2, nil)
		mockOrderRepo.On("CreateOrder", mock.Anything, mock.AnythingOfType("Order.Order")).Return(Order.Order{}, errors.New("failed to create order"))

		result, err := useCase.CreateOrder(ctx, "1")

		assert.Error(t, err)
		assert.Equal(t, "failed to create order", err.Error())
		assert.Equal(t, Order.Order{}, result)
		mockTransactionIDRepo.AssertExpectations(t)
		mockStockRepo.AssertExpectations(t)
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("Test ChangeOrderStatus Success", func(t *testing.T) {
		mockOrderRepo, _, _, useCase := setupOrderUseCase()

		sampleOrder := Order.Order{
			OID:          uuid.New(),
			OTranID:      "1",
			OPrice:       100.00,
			ODestination: "Bangkok",
			OStatus:      "New",
			OPaid:        false,
		}

		updatedOrder := sampleOrder
		updatedOrder.OPaid = true
		updatedOrder.OStatus = "Paid"

		mockOrderRepo.On("GetOrderByID", mock.Anything, sampleOrder.OID.String()).Return(sampleOrder, nil)
		mockOrderRepo.On("ChangeOrderStatus", mock.Anything, updatedOrder, sampleOrder.OID.String()).Return(updatedOrder, nil)

		result, err := useCase.ChangeOrderStatus(ctx, sampleOrder.OID.String(), "Paid")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, updatedOrder, result)
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("Test ChangeOrderStatus Failure", func(t *testing.T) {
		mockOrderRepo, _, _, useCase := setupOrderUseCase()

		sampleOrder := Order.Order{
			OID:          uuid.New(),
			OTranID:      "1",
			OPrice:       100.00,
			ODestination: "Bangkok",
			OStatus:      "New",
			OPaid:        false,
		}

		updatedOrder := sampleOrder
		updatedOrder.OStatus = "Paid"
		updatedOrder.OPaid = true

		mockOrderRepo.On("GetOrderByID", mock.Anything, sampleOrder.OID.String()).Return(sampleOrder, nil)
		mockOrderRepo.On("ChangeOrderStatus", mock.Anything, updatedOrder, sampleOrder.OID.String()).Return(Order.Order{}, errors.New("failed to change status"))

		result, err := useCase.ChangeOrderStatus(ctx, sampleOrder.OID.String(), "Paid")

		assert.Error(t, err)
		assert.Equal(t, "failed to change status", err.Error())
		assert.Equal(t, Order.Order{}, result)
		mockOrderRepo.AssertExpectations(t)
	})

}
