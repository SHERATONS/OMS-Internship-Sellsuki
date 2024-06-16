package Repository

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"gorm.io/gorm"
)

type OrderCalculateRepo struct {
	Db *gorm.DB
}

func (o OrderCalculateRepo) GetAllOrders() ([]Entities.OrderCalculate, error) {
	var TransactionID []Entities.OrderCalculate
	err := o.Db.Order("o_tran_id").Find(&TransactionID).Error
	return TransactionID, err
}

func (o OrderCalculateRepo) GetOrderByTransactionID(transactionID string) (Entities.OrderCalculate, error) {
	var transaction Entities.OrderCalculate
	err := o.Db.Where("o_tran_id = ?", transactionID).First(&transaction).Error
	return transaction, err
}

func (o OrderCalculateRepo) CreateTransactionID(orderCalculate Entities.OrderCalculate) (Entities.OrderCalculate, error) {
	err := o.Db.Create(&orderCalculate).Error
	return orderCalculate, err
}

func (o OrderCalculateRepo) DeleteTransactionID(transactionID string) error {
	err := o.Db.Where("o_tran_id = ?", transactionID).Delete(&Entities.OrderCalculate{}).Error
	return err
}

func NewOrderCalculateRepo(db *gorm.DB) IOrderCalculateRepo {
	return OrderCalculateRepo{Db: db}
}
