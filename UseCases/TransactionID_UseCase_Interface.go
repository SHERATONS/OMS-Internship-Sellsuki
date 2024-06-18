package UseCases

import "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"

type ITransactionIDCase interface {
	GetAllOrders() ([]Entities.TransactionID, error)
	GetOrderByTransactionID(transactionID string) (Entities.TransactionID, error)
	CreateTransactionID(orderCalculate Entities.TransactionID) (Entities.TransactionID, error)
	DeleteTransactionID(transactionID string) error
}
