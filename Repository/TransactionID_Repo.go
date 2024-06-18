package Repository

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Model"
	"gorm.io/gorm"
	"log"
)

type OrderCalculateRepo struct {
	Db *gorm.DB
}

func (o OrderCalculateRepo) GetAllTransactionIDs() ([]Entities.TransactionID, error) {
	var TransactionID []Entities.TransactionID
	err := o.Db.Order("o_tran_id").Find(&TransactionID).Error
	return TransactionID, err
}

func (o OrderCalculateRepo) GetOrderByTransactionID(transactionID string) (Entities.TransactionID, error) {
	var transaction Entities.TransactionID
	err := o.Db.Where("o_tran_id = ?", transactionID).First(&transaction).Error
	return transaction, err
}

func (o OrderCalculateRepo) CreateTransactionID(orderCalculate Entities.TransactionID) (Entities.TransactionID, error) {
	err := o.Db.Create(&orderCalculate).Error
	return orderCalculate, err
}

func (o OrderCalculateRepo) DeleteTransactionID(transactionID string) error {
	err := o.Db.Where("o_tran_id = ?", transactionID).Delete(&Entities.TransactionID{}).Error
	return err
}

func NewOrderCalculateRepo(db *gorm.DB) IOrderCalculateRepo {
	err := db.AutoMigrate(&Model.TransactionID{})
	if err != nil {
		log.Fatalf("Failed to auto migrate TransactionID: %v", err)
	}
	return OrderCalculateRepo{Db: db}
}
