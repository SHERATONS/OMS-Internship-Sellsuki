package Order

import (
	"context"
	"errors"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Order"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Model"
	"gorm.io/gorm"
	"log"
)

type OrderRepo struct {
	Db *gorm.DB
}

func (o *OrderRepo) GetOrderByID(ctx context.Context, orderId string) (Order.Order, error) {
	ctx, span := tracer.Start(ctx, "GetOrderByID_Repo")
	defer span.End()

	var order Order.Order

	err := o.Db.Where("o_id = ?", orderId).First(&order).Error
	if err != nil {
		return order, errors.New("order not found")
	}
	return order, nil
}

func (o *OrderRepo) CreateOrder(ctx context.Context, order Order.Order) (Order.Order, error) {
	ctx, span := tracer.Start(ctx, "CreateOrder_Repo")
	defer span.End()

	err := o.Db.Create(&order).Error
	if err != nil {
		return order, errors.New("failed to create order")
	}
	return order, nil
}

func (o *OrderRepo) ChangeOrderStatus(ctx context.Context, order Order.Order, oid string) (Order.Order, error) {
	ctx, span := tracer.Start(ctx, "ChangeOrderStatus_Repo")
	defer span.End()

	var existingOrder Order.Order

	err := o.Db.First(&existingOrder, "o_id = ?", oid).Error
	if err != nil {
		return Order.Order{}, err
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
