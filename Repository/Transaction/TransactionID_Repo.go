package Transaction

import (
	"context"
	"errors"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/TransactionID"
	"log"

	"github.com/SHERATONS/OMS-Sellsuki-Internship/Model"
	"gorm.io/gorm"
)

type TransactionIDRepo struct {
	Db *gorm.DB
}

func (o TransactionIDRepo) GetAllTransactionIDs(ctx context.Context) ([]TransactionID.TransactionID, error) {
	ctx, span := tracer.Start(ctx, "GetAllTransactionIDs_Repo")
	defer span.End()

	var TransactionID []TransactionID.TransactionID

	err := o.Db.Order("t_id").Find(&TransactionID).Error
	if err != nil {
		return TransactionID, err
	}

	return TransactionID, nil
}

func (o TransactionIDRepo) GetOrderByTransactionID(ctx context.Context, transactionID string) (TransactionID.TransactionID, error) {
	ctx, span := tracer.Start(ctx, "GetOrderByTransactionID_Repo")
	defer span.End()

	var transaction TransactionID.TransactionID

	err := o.Db.Where("t_id = ?", transactionID).First(&transaction).Error
	if err != nil {
		return transaction, errors.New("transaction ID not found")
	}

	return transaction, nil
}

func (o TransactionIDRepo) CreateTransactionID(ctx context.Context, TransactionInfo TransactionID.TransactionID) (TransactionID.TransactionID, error) {
	ctx, span := tracer.Start(ctx, "CreateTransactionID_Repo")
	defer span.End()

	err := o.Db.Create(&TransactionInfo).Error
	if err != nil {
		return TransactionInfo, errors.New("failed to create transaction ID")
	}

	return TransactionInfo, nil
}

func (o TransactionIDRepo) DeleteTransactionID(ctx context.Context, transactionID string) error {
	ctx, span := tracer.Start(ctx, "DeleteTransactionID_Repo")
	defer span.End()

	err := o.Db.Where("t_id = ?", transactionID).Delete(&TransactionID.TransactionID{}).Error
	if err != nil {
		return errors.New("failed to delete transaction ID")
	}

	return err
}

func NewTransactionIDRepo(db *gorm.DB) ITransactionIDRepo {
	err := db.AutoMigrate(&Model.TransactionID{})
	if err != nil {
		log.Fatalf("Failed to auto migrate TransactionID: %v", err)
	}
	return TransactionIDRepo{Db: db}
}
