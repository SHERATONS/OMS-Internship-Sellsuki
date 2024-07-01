package Transaction

import (
	"context"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
)

type ITransactionIDRepo interface {
	GetAllTransactionIDs(ctx context.Context) ([]Entities.TransactionID, error)
	GetOrderByTransactionID(ctx context.Context, transactionID string) (Entities.TransactionID, error)
	CreateTransactionID(ctx context.Context, transactionInfo Entities.TransactionID) (Entities.TransactionID, error)
	DeleteTransactionID(ctx context.Context, transactionID string) error
}
