package Repository

import "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"

type ITransactionIDRepo interface {
	GetAllTransactionIDs() ([]Entities.TransactionID, error)
	GetOrderByTransactionID(transactionID string) (Entities.TransactionID, error)
	CreateTransactionID(transactionInfo Entities.TransactionID) (Entities.TransactionID, error)
	DeleteTransactionID(transactionID string) error
}
