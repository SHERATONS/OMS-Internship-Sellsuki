package UseCases

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository"
)

type TransactionIDUseCase struct {
	Repo Repository.ITransactionIDRepo
}

func (o TransactionIDUseCase) GetAllTransactionIDs() ([]Entities.TransactionID, error) {
	return o.Repo.GetAllTransactionIDs()
}

func (o TransactionIDUseCase) GetOrderByTransactionID(transactionID string) (Entities.TransactionID, error) {
	return o.Repo.GetOrderByTransactionID(transactionID)
}

func (o TransactionIDUseCase) CreateTransactionID(transactionInfo Entities.TransactionID) (Entities.TransactionID, error) {
	return o.Repo.CreateTransactionID(transactionInfo)
}

func (o TransactionIDUseCase) DeleteTransactionID(transactionID string) error {
	return o.Repo.DeleteTransactionID(transactionID)
}

func NewTransactionIDUseCase(repo Repository.ITransactionIDRepo) ITransactionIDUseCase {
	return TransactionIDUseCase{Repo: repo}
}
