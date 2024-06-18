package UseCases

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository"
)

type OrderCalculateUseCases struct {
	Repo Repository.IOrderCalculateRepo
}

func (o OrderCalculateUseCases) GetAllOrders() ([]Entities.TransactionID, error) {
	return o.Repo.GetAllTransactionIDs()
}

func (o OrderCalculateUseCases) GetOrderByTransactionID(transactionID string) (Entities.TransactionID, error) {
	return o.Repo.GetOrderByTransactionID(transactionID)
}

func (o OrderCalculateUseCases) CreateTransactionID(orderCalculate Entities.TransactionID) (Entities.TransactionID, error) {
	return o.Repo.CreateTransactionID(orderCalculate)
}

func (o OrderCalculateUseCases) DeleteTransactionID(transactionID string) error {
	return o.Repo.DeleteTransactionID(transactionID)
}

func NewOrderCalculateUseCases(repo Repository.IOrderCalculateRepo) ITransactionIDCase {
	return OrderCalculateUseCases{Repo: repo}
}
