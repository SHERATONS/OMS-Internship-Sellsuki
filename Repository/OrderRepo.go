package Repository

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"gorm.io/gorm"
)

type OrderRepo struct {
	Db *gorm.DB
}

func (o *OrderRepo) CreateOrder(order Entities.Order) (Entities.Order, error) {
	err := o.Db.Create(&order).Error
	return order, err
}

func (o *OrderRepo) ChangeOrderStatus(order Entities.Order, oid int64) (Entities.Order, error) {
	err := o.Db.Where("oid = ?", oid).Save(&order).Error
	return order, err
}

func NewOrderRepo(db *gorm.DB) IOrderRepo {
	return &OrderRepo{Db: db}
}
