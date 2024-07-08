package MockRepository

import (
	"context"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Order"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepo struct {
	mock.Mock
}

func (m *MockOrderRepo) GetOrderByID(ctx context.Context, orderId string) (Order.Order, error) {
	args := m.Called(ctx, orderId)
	return args.Get(0).(Order.Order), args.Error(1)
}

func (m *MockOrderRepo) CreateOrder(ctx context.Context, order Order.Order) (Order.Order, error) {
	args := m.Called(ctx, order)
	return args.Get(0).(Order.Order), args.Error(1)
}

func (m *MockOrderRepo) ChangeOrderStatus(ctx context.Context, order Order.Order, oid string) (Order.Order, error) {
	args := m.Called(ctx, order, oid)
	return args.Get(0).(Order.Order), args.Error(1)
}
