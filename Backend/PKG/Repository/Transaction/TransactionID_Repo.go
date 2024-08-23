package Transaction

import (
	"context"
	"errors"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/TransactionID"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Model"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"log"
	"reflect"

	"gorm.io/gorm"
)

type TransactionIDRepo struct {
	Db *gorm.DB
}

func (o *TransactionIDRepo) GetAllTransactionIDs(ctx context.Context) ([]TransactionID.TransactionID, error) {
	ctx, span := tracer.Start(ctx, "GetAllTransactionIDs_Repo")
	defer span.End()

	var transactionID []TransactionID.TransactionID

	err := o.Db.Order("t_id").Find(&transactionID).Error
	if err != nil {
		return transactionID, err
	}

	return transactionID, nil
}

func (o *TransactionIDRepo) GetOrderByTransactionID(ctx context.Context, transactionID string) (TransactionID.TransactionID, error) {
	ctx, span := tracer.Start(ctx, "GetOrderByTransactionID_Repo")
	defer span.End()

	var transaction TransactionID.TransactionID

	err := o.Db.Where("t_id = ?", transactionID).First(&transaction).Error
	if err != nil {
		return transaction, errors.New("transaction ID not found")
	}

	return transaction, nil
}

func (o *TransactionIDRepo) CreateTransactionID(ctx context.Context, transactionInfo TransactionID.TransactionID) (TransactionID.TransactionID, error) {
	ctx, span := tracer.Start(ctx, "CreateTransactionID_Repo")
	defer span.End()

	err := o.Db.Create(&transactionInfo).Error
	if err != nil {
		return transactionInfo, errors.New("failed to create transaction ID")
	}

	return transactionInfo, nil
}

func (o *TransactionIDRepo) DeleteTransactionID(ctx context.Context, transactionID string) error {
	ctx, span := tracer.Start(ctx, "DeleteTransactionID_Repo")
	defer span.End()

	err := o.Db.Where("t_id = ?", transactionID).Delete(&TransactionID.TransactionID{}).Error
	if err != nil {
		return errors.New("failed to delete transaction ID")
	}

	return err
}

func (o *TransactionIDRepo) SetTransactionSubAttributes(transactionData any, sp trace.Span) {
	if transaction, ok := transactionData.(TransactionID.TransactionID); ok {
		sp.SetAttributes(
			attribute.String("TransactionID", transaction.TID),
			attribute.Float64("TransactionPrice", transaction.TPrice),
			attribute.String("TransactionDestination", transaction.TDestination),
			attribute.String("TransactionProductList", transaction.TProductList),
		)
	} else if transactions, ok := transactionData.(*[]TransactionID.TransactionID); ok {
		transactionIDs := make([]string, len(*transactions))
		transactionPrices := make([]float64, len(*transactions))
		transactionDestinations := make([]string, len(*transactions))
		transactionProductLists := make([]string, len(*transactions))

		for _, transaction := range *transactions {
			transactionIDs = append(transactionIDs, transaction.TID)
			transactionPrices = append(transactionPrices, transaction.TPrice)
			transactionDestinations = append(transactionDestinations, transaction.TDestination)
			transactionProductLists = append(transactionProductLists, transaction.TProductList)
		}

		sp.SetAttributes(
			attribute.StringSlice("TransactionID", transactionIDs),
			attribute.Float64Slice("TransactionPrice", transactionPrices),
			attribute.StringSlice("TransactionDestination", transactionDestinations),
			attribute.StringSlice("TransactionProductList", transactionProductLists),
		)
	} else {
		sp.RecordError(errors.New("invalid type: " + reflect.TypeOf(transactionData).String()))
	}
}

func NewTransactionIDRepo(db *gorm.DB) ITransactionIDRepo {
	err := db.AutoMigrate(&Model.TransactionID{})
	if err != nil {
		log.Fatalf("Failed to auto migrate TransactionID: %v", err)
	}
	return &TransactionIDRepo{Db: db}
}
