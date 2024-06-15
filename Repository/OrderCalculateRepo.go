package Repository

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"gorm.io/gorm"
)

type OrderCalculateRepo struct {
	Db *gorm.DB
}

//func (o OrderCalculateRepo) GetAllOrders() ([]Entities.OrderCalculate, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (o OrderCalculateRepo) GetOrderByTransactionID(transactionID string) (Entities.OrderCalculate, error) {
//	//TODO implement me
//	panic("implement me")
//}

func (o OrderCalculateRepo) CreateTransactionID(orderCalculate Entities.OrderCalculate) (Entities.OrderCalculate, error) {
	err := o.Db.Create(&orderCalculate).Error
	return orderCalculate, err
}

//func (o OrderCalculateRepo) DeleteTransactionID(transactionID string) error {
//	//TODO implement me
//	panic("implement me")
//}

func NewOrderCalculateRepo(db *gorm.DB) IOrderCalculateRepo {
	return OrderCalculateRepo{Db: db}
}
