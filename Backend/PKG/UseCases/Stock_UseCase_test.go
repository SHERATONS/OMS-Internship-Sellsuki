package UseCases

import (
	"context"
	"errors"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Product"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Stock"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/PKG/UseCases/MockRepository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func setupStockUseCase() (*MockRepository.MockStockRepo, *MockRepository.MockProductRepo, IStockUseCase) {
	mockStockRepo := new(MockRepository.MockStockRepo)
	mockProductRepo := new(MockRepository.MockProductRepo)
	useCases := NewStockUseCase(mockStockRepo, mockProductRepo)
	return mockStockRepo, mockProductRepo, useCases
}

func TestStockUseCase(t *testing.T) {
	ctx := context.TODO()

	t.Run("Test GetAllStocks Success", func(t *testing.T) {
		mockStockRepo, _, useCases := setupStockUseCase()

		sampleStocks := []Stock.Stock{
			{SID: "1", SQuantity: 10, SUpdated: time.Now()},
			{SID: "2", SQuantity: 10, SUpdated: time.Now()},
		}

		mockStockRepo.On("GetAllStocks", mock.Anything).Return(sampleStocks, nil)

		result, err := useCases.GetAllStocks(ctx)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, sampleStocks, result)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("Test GetAllStocks Failure", func(t *testing.T) {
		mockStockRepo, _, useCases := setupStockUseCase()

		mockStockRepo.On("GetAllStocks", mock.Anything).Return([]Stock.Stock{}, assert.AnError)

		result, err := useCases.GetAllStocks(ctx)

		assert.Error(t, err)
		assert.Equal(t, []Stock.Stock{}, result)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("Test GetStockByID Success", func(t *testing.T) {
		mockStockRepo, _, useCases := setupStockUseCase()

		sampleStock := Stock.Stock{
			SID:       "1",
			SQuantity: 10,
			SUpdated:  time.Now(),
		}

		mockStockRepo.On("GetStockByID", mock.Anything, "1").Return(sampleStock, nil)

		result, err := useCases.GetStockByID(ctx, "1")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, sampleStock, result)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("Test GetStockByID Failure", func(t *testing.T) {
		mockStockRepo, _, useCases := setupStockUseCase()

		mockStockRepo.On("GetStockByID", mock.Anything, "1").Return(Stock.Stock{}, errors.New("stock Not Found"))

		result, err := useCases.GetStockByID(ctx, "1")

		assert.Error(t, err)
		assert.Equal(t, "stock Not Found", err.Error())
		assert.Equal(t, Stock.Stock{}, result)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("Test CreateStock Success", func(t *testing.T) {
		mockStockRepo, mockProductRepo, useCases := setupStockUseCase()

		sampleStock := Stock.Stock{
			SID:       "1",
			SQuantity: 10,
			SUpdated:  time.Now(),
		}

		mockProductRepo.On("GetProductByID", mock.Anything, "1").Return(Product.Product{}, nil)
		mockStockRepo.On("CreateStock", mock.Anything, sampleStock).Return(sampleStock, nil)

		result, err := useCases.CreateStock(ctx, sampleStock)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, sampleStock, result)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("Test CreateStock Failure", func(t *testing.T) {
		mockStockRepo, mockProductRepo, useCases := setupStockUseCase()

		sampleStock := Stock.Stock{
			SID:       "1",
			SQuantity: 10,
			SUpdated:  time.Now(),
		}

		mockProductRepo.On("GetProductByID", mock.Anything, "1").Return(Product.Product{}, nil)
		mockStockRepo.On("CreateStock", mock.Anything, sampleStock).Return(Stock.Stock{}, errors.New("failed to update stock"))

		result, err := useCases.CreateStock(ctx, sampleStock)

		assert.Error(t, err)
		assert.Equal(t, "failed to update stock", err.Error())
		assert.Equal(t, Stock.Stock{}, result)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("Test UpdateStock Success", func(t *testing.T) {
		mockStockRepo, mockProductRepo, useCases := setupStockUseCase()

		sampleStock := Stock.Stock{
			SID:       "1",
			SQuantity: 10,
			SUpdated:  time.Now(),
		}

		mockStockRepo.On("GetStockByID", mock.Anything, "1").Return(sampleStock, nil)
		mockProductRepo.On("GetProductByID", mock.Anything, "1").Return(Product.Product{}, nil)
		mockStockRepo.On("UpdateStock", mock.Anything, sampleStock, "1").Return(sampleStock, nil)

		result, err := useCases.UpdateStock(ctx, sampleStock, "1")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, sampleStock, result)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("Test UpdateStock Failure", func(t *testing.T) {
		mockStockRepo, mockProductRepo, useCases := setupStockUseCase()

		sampleStock := Stock.Stock{
			SID:       "1",
			SQuantity: 10,
			SUpdated:  time.Now(),
		}

		mockStockRepo.On("GetStockByID", mock.Anything, "1").Return(sampleStock, nil)
		mockProductRepo.On("GetProductByID", mock.Anything, "1").Return(Product.Product{}, nil)
		mockStockRepo.On("UpdateStock", mock.Anything, sampleStock, "1").Return(Stock.Stock{}, errors.New("failed to delete stock"))

		result, err := useCases.UpdateStock(ctx, sampleStock, "1")

		assert.Error(t, err)
		assert.Equal(t, "failed to delete stock", err.Error())
		assert.Equal(t, Stock.Stock{}, result)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("Test DeleteStock Success", func(t *testing.T) {
		mockStockRepo, _, useCases := setupStockUseCase()

		mockStockRepo.On("GetStockByID", mock.Anything, "1").Return(Stock.Stock{}, nil)
		mockStockRepo.On("DeleteStock", mock.Anything, "1").Return(nil)

		err := useCases.DeleteStock(ctx, "1")

		assert.NoError(t, err)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("Test DeleteStock Failure", func(t *testing.T) {
		mockStockRepo, _, useCases := setupStockUseCase()

		mockStockRepo.On("GetStockByID", mock.Anything, "1").Return(Stock.Stock{}, nil)
		mockStockRepo.On("DeleteStock", mock.Anything, "1").Return(errors.New("failed to delete stock"))

		err := useCases.DeleteStock(ctx, "1")

		assert.Error(t, err)
		mockStockRepo.AssertExpectations(t)
	})
}
