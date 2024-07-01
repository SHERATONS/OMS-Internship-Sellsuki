package Order

import (
	"context"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
)

type IOrderRepo interface {
	CreateOrder(ctx context.Context, order Entities.Order) (Entities.Order, error)
	ChangeOrderStatus(ctx context.Context, order Entities.Order, orderID string) (Entities.Order, error)
	GetOrderByID(ctx context.Context, orderID string) (Entities.Order, error)
}
