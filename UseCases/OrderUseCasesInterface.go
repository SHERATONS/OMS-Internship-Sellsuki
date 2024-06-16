package UseCases

import "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"

type IOrderCase interface {
	CreateOrder(order Entities.Order) (Entities.Order, error)
	ChangeOrderStatus(order Entities.Order, oid int64) (Entities.Order, error)
	GetOrderById(orderId int64) (Entities.Order, error)
}
