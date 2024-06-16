package UseCases

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository"
)

type OrderCalculateUseCases struct {
	Repo Repository.IOrderCalculateRepo
}

func (o OrderCalculateUseCases) GetAllOrders() ([]Entities.OrderCalculate, error) {
	return o.Repo.GetAllOrders()
}

func (o OrderCalculateUseCases) GetOrderByTransactionID(transactionID string) (Entities.OrderCalculate, error) {
	return o.Repo.GetOrderByTransactionID(transactionID)
}

func (o OrderCalculateUseCases) CreateTransactionID(orderCalculate Entities.OrderCalculate) (Entities.OrderCalculate, error) {
	return o.Repo.CreateTransactionID(orderCalculate)
}

func (o OrderCalculateUseCases) DeleteTransactionID(transactionID string) error {
	return o.Repo.DeleteTransactionID(transactionID)
}

func NewOrderCalculateUseCases(repo Repository.IOrderCalculateRepo) IOrderCalculateCase {
	return OrderCalculateUseCases{Repo: repo}
}
