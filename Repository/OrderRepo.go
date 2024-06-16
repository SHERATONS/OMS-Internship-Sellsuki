package Repository

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"gorm.io/gorm"
)

type OrderRepo struct {
	Db *gorm.DB
}

func (o *OrderRepo) GetOrderById(orderId int64) (Entities.Order, error) {
	var order Entities.Order
	err := o.Db.Where("o_id = ?", orderId).First(&order).Error
	return order, err
}

func (o *OrderRepo) CreateOrder(order Entities.Order) (Entities.Order, error) {
	err := o.Db.Omit("o_id").Create(&order).Error
	return order, err
}

func (o *OrderRepo) ChangeOrderStatus(order Entities.Order, oid int64) (Entities.Order, error) {
	var existingOrder Entities.Order
	err := o.Db.First(&existingOrder, "o_id = ?", oid).Error
	if err != nil {
		return Entities.Order{}, err
	}

	existingOrder.OStatus = order.OStatus
	existingOrder.OPaid = order.OPaid

	err = o.Db.Where("o_id = ?", oid).Model(&existingOrder).Updates(map[string]interface{}{
		"o_status": existingOrder.OStatus,
		"o_paid":   existingOrder.OPaid,
	}).Error

	return existingOrder, err
}

func NewOrderRepo(db *gorm.DB) IOrderRepo {
	return &OrderRepo{Db: db}
}
