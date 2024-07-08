package UseCases

import (
	"context"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Order"
	"go.opentelemetry.io/otel"
)

type IOrderUseCase interface {
	CreateOrder(ctx context.Context, transactionID string) (Order.Order, error)
	ChangeOrderStatus(ctx context.Context, oid string, oStatus string) (Order.Order, error)
	GetOrderById(ctx context.Context, orderId string) (Order.Order, error)
}

var tracerOrder = otel.Tracer("Order_UseCase")
