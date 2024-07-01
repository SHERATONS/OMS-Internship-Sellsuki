package Transaction

import (
	"context"
	"errors"
	"log"

	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Model"
	"gorm.io/gorm"
)

type TransactionIDRepo struct {
	Db *gorm.DB
}

func (o TransactionIDRepo) GetAllTransactionIDs(ctx context.Context) ([]Entities.TransactionID, error) {
	var TransactionID []Entities.TransactionID

	err := o.Db.Order("t_id").Find(&TransactionID).Error
	if err != nil {
		return TransactionID, err
	}

	return TransactionID, nil
}

func (o TransactionIDRepo) GetOrderByTransactionID(ctx context.Context, transactionID string) (Entities.TransactionID, error) {
	var transaction Entities.TransactionID

	err := o.Db.Where("t_id = ?", transactionID).First(&transaction).Error
	if err != nil {
		return transaction, errors.New("transaction ID not found")
	}

	return transaction, nil
}

func (o TransactionIDRepo) CreateTransactionID(ctx context.Context, TransactionInfo Entities.TransactionID) (Entities.TransactionID, error) {
	err := o.Db.Create(&TransactionInfo).Error
	if err != nil {
		return TransactionInfo, errors.New("failed to create transaction ID")
	}

	return TransactionInfo, nil
}

func (o TransactionIDRepo) DeleteTransactionID(ctx context.Context, transactionID string) error {
	err := o.Db.Where("t_id = ?", transactionID).Delete(&Entities.TransactionID{}).Error
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
