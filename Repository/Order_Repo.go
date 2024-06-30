package Repository

import (
	"errors"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Model"
	"gorm.io/gorm"
	"log"
)

type OrderRepo struct {
	Db *gorm.DB
}

func (o *OrderRepo) GetOrderByID(orderId string) (Entities.Order, error) {
	var order Entities.Order

	err := o.Db.Where("o_id = ?", orderId).First(&order).Error
	if err != nil {
		return order, errors.New("order not found")
	}
	return order, nil
}

func (o *OrderRepo) CreateOrder(order Entities.Order) (Entities.Order, error) {
	err := o.Db.Create(&order).Error
	if err != nil {
		return order, errors.New("failed to create order")
	}
	return order, nil
}

func (o *OrderRepo) ChangeOrderStatus(order Entities.Order, oid string) (Entities.Order, error) {
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
	err := db.AutoMigrate(&Model.Order{})
	if err != nil {
		log.Fatalf("Failed to auto migrate Order: %v", err)
	}
	return &OrderRepo{Db: db}
}
