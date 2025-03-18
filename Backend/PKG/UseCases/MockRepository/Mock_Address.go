package MockRepository

import (
	"context"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Address"
	"github.com/stretchr/testify/mock"
)

type MockAddressRepo struct {
	mock.Mock
}

func (m *MockAddressRepo) GetAddressByCity(ctx context.Context, city string) (Address.Address, error) {
	args := m.Called(ctx, city)
	return args.Get(0).(Address.Address), args.Error(1)
}

func (m *MockAddressRepo) CreateAddress(ctx context.Context, address Address.Address) (Address.Address, error) {
	args := m.Called(ctx, address)
	return args.Get(0).(Address.Address), args.Error(1)
}

func (m *MockAddressRepo) UpdateAddress(ctx context.Context, address Address.Address, city string) (Address.Address, error) {
	args := m.Called(ctx, address, city)
	return args.Get(0).(Address.Address), args.Error(1)
}

func (m *MockAddressRepo) DeleteAddress(ctx context.Context, city string) error {
	args := m.Called(ctx, city)
	return args.Error(0)
}
