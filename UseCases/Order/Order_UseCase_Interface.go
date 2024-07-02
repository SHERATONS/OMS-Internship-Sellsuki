package Order

import (
	"context"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Order"
	"go.opentelemetry.io/otel"
)

type IOrderUseCase interface {
	CreateOrder(ctx context.Context, TransactionID string) (Order.Order, error)
	ChangeOrderStatus(ctx context.Context, oid string, oStatus string) (Order.Order, error)
	GetOrderById(ctx context.Context, orderId string) (Order.Order, error)
}

var tracer = otel.Tracer("Order_UseCase")
