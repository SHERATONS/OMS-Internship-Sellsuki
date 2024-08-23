package MockRepository

import (
	"context"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Backend/Entities/Stock"
	"github.com/stretchr/testify/mock"
)

type MockStockRepo struct {
	mock.Mock
}

func (m *MockStockRepo) GetAllStocks(ctx context.Context) ([]Stock.Stock, error) {
	args := m.Called(ctx)
	return args.Get(0).([]Stock.Stock), args.Error(1)
}

func (m *MockStockRepo) GetStockByID(ctx context.Context, stockID string) (Stock.Stock, error) {
	args := m.Called(ctx, stockID)
	return args.Get(0).(Stock.Stock), args.Error(1)
}

func (m *MockStockRepo) CreateStock(ctx context.Context, stock Stock.Stock) (Stock.Stock, error) {
	args := m.Called(ctx, stock)
	return args.Get(0).(Stock.Stock), args.Error(1)
}

func (m *MockStockRepo) UpdateStock(ctx context.Context, stock Stock.Stock, stockID string) (Stock.Stock, error) {
	args := m.Called(ctx, stock, stockID)
	return args.Get(0).(Stock.Stock), args.Error(1)
}

func (m *MockStockRepo) DeleteStock(ctx context.Context, stockID string) error {
	args := m.Called(ctx, stockID)
	return args.Error(0)
}
