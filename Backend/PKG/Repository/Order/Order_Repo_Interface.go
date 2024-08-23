package Order

import (
	"context"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Order"
	"go.opentelemetry.io/otel"
)

type IOrderRepo interface {
	CreateOrder(ctx context.Context, order Order.Order) (Order.Order, error)
	ChangeOrderStatus(ctx context.Context, order Order.Order, orderID string) (Order.Order, error)
	GetOrderByID(ctx context.Context, orderID string) (Order.Order, error)
}

var tracer = otel.Tracer("Order_Repo")
