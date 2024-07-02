package Transaction

import (
	"context"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/TransactionID"
	"go.opentelemetry.io/otel"
)

type ITransactionIDRepo interface {
	GetAllTransactionIDs(ctx context.Context) ([]TransactionID.TransactionID, error)
	GetOrderByTransactionID(ctx context.Context, transactionID string) (TransactionID.TransactionID, error)
	CreateTransactionID(ctx context.Context, transactionInfo TransactionID.TransactionID) (TransactionID.TransactionID, error)
	DeleteTransactionID(ctx context.Context, transactionID string) error
}

var tracer = otel.Tracer("TransactionID_Repo")
