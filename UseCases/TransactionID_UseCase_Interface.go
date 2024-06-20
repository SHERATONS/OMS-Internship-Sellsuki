package UseCases

import "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"

type ITransactionIDUseCase interface {
	GetAllTransactionIDs() ([]Entities.TransactionID, error)
	GetOrderByTransactionID(transactionID string) (Entities.TransactionID, error)
	CreateTransactionID(transactionInfo Entities.TransactionID) (Entities.TransactionID, error)
	DeleteTransactionID(transactionID string) error
}
