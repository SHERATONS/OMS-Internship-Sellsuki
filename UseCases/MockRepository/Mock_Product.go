package MockRepository

import (
	"context"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Product"
	"github.com/stretchr/testify/mock"
)

type MockProductRepo struct {
	mock.Mock
}

func (m *MockProductRepo) GetAllProducts(ctx context.Context) ([]Product.Product, error) {
	args := m.Called(ctx)
	return args.Get(0).([]Product.Product), args.Error(1)
}

func (m *MockProductRepo) GetProductByID(ctx context.Context, productID string) (Product.Product, error) {
	args := m.Called(ctx, productID)
	return args.Get(0).(Product.Product), args.Error(1)
}

func (m *MockProductRepo) CreateProduct(ctx context.Context, product Product.Product) (Product.Product, error) {
	args := m.Called(ctx, product)
	return args.Get(0).(Product.Product), args.Error(1)
}

func (m *MockProductRepo) UpdateProduct(ctx context.Context, product Product.Product, productID string) (Product.Product, error) {
	args := m.Called(ctx, product, productID)
	return args.Get(0).(Product.Product), args.Error(1)
}

func (m *MockProductRepo) DeleteProduct(ctx context.Context, productID string) error {
	args := m.Called(ctx, productID)
	return args.Error(0)
}
