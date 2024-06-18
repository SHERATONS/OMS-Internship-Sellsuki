package Repository

import "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"

type IOrderCalculateRepo interface {
	GetAllTransactionIDs() ([]Entities.TransactionID, error)
	GetOrderByTransactionID(transactionID string) (Entities.TransactionID, error)
	CreateTransactionID(transaction Entities.TransactionID) (Entities.TransactionID, error)
	DeleteTransactionID(transactionID string) error
}
