package Repository

import (
	"log"

	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Model"
	"gorm.io/gorm"
)

type TransactionIDRepo struct {
	Db *gorm.DB
}

func (o TransactionIDRepo) GetAllTransactionIDs() ([]Entities.TransactionID, error) {
	var TransactionID []Entities.TransactionID
	err := o.Db.Order("t_id").Find(&TransactionID).Error
	return TransactionID, err
}

func (o TransactionIDRepo) GetOrderByTransactionID(transactionID string) (Entities.TransactionID, error) {
	var transaction Entities.TransactionID
	err := o.Db.Where("t_id = ?", transactionID).First(&transaction).Error
	return transaction, err
}

func (o TransactionIDRepo) CreateTransactionID(TransactionInfo Entities.TransactionID) (Entities.TransactionID, error) {
	err := o.Db.Create(&TransactionInfo).Error
	return TransactionInfo, err
}

func (o TransactionIDRepo) DeleteTransactionID(transactionID string) error {
	err := o.Db.Where("t_id = ?", transactionID).Delete(&Entities.TransactionID{}).Error
	return err
}

func NewTransactionIDRepo(db *gorm.DB) ITransactionIDRepo {
	err := db.AutoMigrate(&Model.TransactionID{})
	if err != nil {
		log.Fatalf("Failed to auto migrate TransactionID: %v", err)
	}
	return TransactionIDRepo{Db: db}
}
