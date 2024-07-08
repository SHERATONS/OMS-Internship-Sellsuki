package MockRepository

import (
	"context"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/TransactionID"
	"github.com/stretchr/testify/mock"
)

type MockTransactionIDRepo struct {
	mock.Mock
}

func (m *MockTransactionIDRepo) GetAllTransactionIDs(ctx context.Context) ([]TransactionID.TransactionID, error) {
	args := m.Called(ctx)
	return args.Get(0).([]TransactionID.TransactionID), args.Error(1)
}

func (m *MockTransactionIDRepo) GetOrderByTransactionID(ctx context.Context, transactionID string) (TransactionID.TransactionID, error) {
	args := m.Called(ctx, transactionID)
	return args.Get(0).(TransactionID.TransactionID), args.Error(1)
}

func (m *MockTransactionIDRepo) CreateTransactionID(ctx context.Context, transactionInfo TransactionID.TransactionID) (TransactionID.TransactionID, error) {
	args := m.Called(ctx, transactionInfo)
	return args.Get(0).(TransactionID.TransactionID), args.Error(1)
}

func (m *MockTransactionIDRepo) DeleteTransactionID(ctx context.Context, transactionID string) error {
	args := m.Called(ctx, transactionID)
	return args.Error(0)
}
