package UseCases

import (
	"context"
	"errors"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Stock"
	"testing"
	"time"

	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Product"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases/MockRepository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupProductUseCase() (*MockRepository.MockProductRepo, *MockRepository.MockStockRepo, IProductUseCase) {
	mockProductRepo := new(MockRepository.MockProductRepo)
	mockStockRepo := new(MockRepository.MockStockRepo)
	useCase := NewProductUseCase(mockProductRepo, mockStockRepo)

	return mockProductRepo, mockStockRepo, useCase
}

func TestProductUseCase(t *testing.T) {
	ctx := context.TODO()

	t.Run("Test GetProductByID Success", func(t *testing.T) {
		mockProductRepo, _, useCase := setupProductUseCase()

		sampleProduct := Product.Product{
			PID:      "1",
			PName:    "Sample Product",
			PPrice:   100.0,
			PDesc:    "Sample Description",
			PCreated: time.Now(),
			PUpdated: time.Now(),
		}

		mockProductRepo.On("GetProductByID", mock.Anything, "1").Return(sampleProduct, nil)

		result, err := useCase.GetProductById(ctx, "1")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, sampleProduct, result)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("Test GetProductByID Failure", func(t *testing.T) {
		mockProductRepo, _, useCase := setupProductUseCase()

		mockProductRepo.On("GetProductByID", mock.Anything, "1").Return(Product.Product{}, errors.New("product not found"))

		result, err := useCase.GetProductById(ctx, "1")

		assert.Error(t, err)
		assert.Equal(t, "product not found", err.Error())
		assert.Equal(t, Product.Product{}, result)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("Test CreateProduct Success", func(t *testing.T) {
		mockProductRepo, _, useCase := setupProductUseCase()

		sampleProduct := Product.Product{
			PID:      "1",
			PName:    "Sample Product",
			PPrice:   100.0,
			PDesc:    "Sample Description",
			PCreated: time.Now(),
			PUpdated: time.Now(),
		}

		mockProductRepo.On("CreateProduct", mock.Anything, sampleProduct).Return(sampleProduct, nil)

		result, err := useCase.CreateProduct(ctx, sampleProduct)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, sampleProduct, result)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("Test CreateProduct Failure", func(t *testing.T) {
		mockProductRepo, _, useCase := setupProductUseCase()

		sampleProduct := Product.Product{
			PID:    "1",
			PName:  "Sample Product",
			PPrice: 100.0,
			PDesc:  "Sample Description",
		}

		mockProductRepo.On("CreateProduct", mock.Anything, sampleProduct).Return(Product.Product{}, errors.New("failed to create product"))

		result, err := useCase.CreateProduct(ctx, sampleProduct)

		assert.Error(t, err)
		assert.Equal(t, "failed to create product", err.Error())
		assert.Equal(t, Product.Product{}, result)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("Test UpdateProduct Success", func(t *testing.T) {
		mockProductRepo, _, useCase := setupProductUseCase()

		updatedProduct := Product.Product{
			PID:      "1",
			PName:    "Updated Product",
			PPrice:   150.0,
			PDesc:    "Updated Description",
			PUpdated: time.Now(),
		}

		mockProductRepo.On("GetProductByID", mock.Anything, "1").Return(updatedProduct, nil)
		mockProductRepo.On("UpdateProduct", mock.Anything, updatedProduct, "1").Return(updatedProduct, nil)

		result, err := useCase.UpdateProduct(ctx, updatedProduct, "1")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, updatedProduct, result)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("Test UpdateProduct Failure", func(t *testing.T) {
		mockProductRepo, _, useCase := setupProductUseCase()

		updatedProduct := Product.Product{
			PID:    "1",
			PName:  "Updated Product",
			PPrice: 150.0,
			PDesc:  "Updated Description",
		}

		mockProductRepo.On("GetProductByID", mock.Anything, "1").Return(updatedProduct, nil)
		mockProductRepo.On("UpdateProduct", mock.Anything, updatedProduct, "1").Return(Product.Product{}, errors.New("failed to Update Product"))

		result, err := useCase.UpdateProduct(ctx, updatedProduct, "1")

		assert.Error(t, err)
		assert.Equal(t, "failed to Update Product", err.Error())
		assert.Equal(t, Product.Product{}, result)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("Test DeleteProductByID Success", func(t *testing.T) {
		mockProductRepo, mockStockRepo, useCase := setupProductUseCase()

		mockProductRepo.On("GetProductByID", mock.Anything, "1").Return(Product.Product{}, nil)
		mockStockRepo.On("GetStockByID", mock.Anything, "1").Return(Stock.Stock{}, nil)
		mockStockRepo.On("DeleteStock", mock.Anything, "1").Return(nil)
		mockProductRepo.On("DeleteProduct", mock.Anything, "1").Return(nil)

		err := useCase.DeleteProductById(ctx, "1")

		assert.NoError(t, err)
		mockProductRepo.AssertExpectations(t)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("Test DeleteProductByID Failure", func(t *testing.T) {
		mockProductRepo, mockStockRepo, useCase := setupProductUseCase()

		mockProductRepo.On("GetProductByID", mock.Anything, "1").Return(Product.Product{}, assert.AnError)

		err := useCase.DeleteProductById(ctx, "1")

		assert.Error(t, err)
		mockProductRepo.AssertExpectations(t)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("Test GetAllProducts Success", func(t *testing.T) {
		mockProductRepo, _, useCase := setupProductUseCase()

		sampleProducts := []Product.Product{
			{PID: "1", PName: "Sample Product 1", PPrice: 100.0, PDesc: "Sample Description 1", PCreated: time.Now(), PUpdated: time.Now()},
			{PID: "2", PName: "Sample Product 2", PPrice: 200.0, PDesc: "Sample Description 2", PCreated: time.Now(), PUpdated: time.Now()},
		}

		mockProductRepo.On("GetAllProducts", mock.Anything).Return(sampleProducts, nil)

		result, err := useCase.GetAllProducts(ctx)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, sampleProducts, result)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("Test GetAllProducts Failure", func(t *testing.T) {
		mockProductRepo, _, useCase := setupProductUseCase()

		mockProductRepo.On("GetAllProducts", mock.Anything).Return([]Product.Product{}, assert.AnError)

		result, err := useCase.GetAllProducts(ctx)

		assert.Error(t, err)
		assert.Equal(t, []Product.Product{}, result)
		mockProductRepo.AssertExpectations(t)
	})
}
