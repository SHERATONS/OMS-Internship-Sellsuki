package Order

import (
	"context"
	"errors"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Order"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Model"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"log"
	"reflect"
	"time"
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

	o.SetOrderSubAttributes(order, span)

	return order, nil
}

func (o *OrderRepo) CreateOrder(ctx context.Context, order Order.Order) (Order.Order, error) {
	ctx, span := tracer.Start(ctx, "CreateOrder_Repo")
	defer span.End()

	order.OCreated = time.Now()

	err := o.Db.Create(&order).Error
	if err != nil {
		return order, errors.New("failed to create order")
	}

	o.SetOrderSubAttributes(order, span)

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

	o.SetOrderSubAttributes(existingOrder, span)

	return existingOrder, err
}

func (o *OrderRepo) SetOrderSubAttributes(orderData any, sp trace.Span) {
	if order, ok := orderData.(Order.Order); ok {
		sp.SetAttributes(
			attribute.String("OrderOID", order.OID.String()),
			attribute.String("OrderOTranID", order.OTranID),
			attribute.Float64("OrderOPrice", order.OPrice),
			attribute.String("OrderODestination", order.ODestination),
			attribute.String("OrderOStatus", order.OStatus),
			attribute.Bool("OrderOPaid", order.OPaid),
			attribute.String("OrderOCreated", order.OCreated.Format(time.RFC3339)),
		)
	} else {
		sp.RecordError(errors.New("invalid type: " + reflect.TypeOf(orderData).String()))
	}
}

func NewOrderRepo(db *gorm.DB) IOrderRepo {
	err := db.AutoMigrate(&Model.Order{})
	if err != nil {
		log.Fatalf("Failed to auto migrate Order: %v", err)
	}
	return &OrderRepo{Db: db}
}
