package Repository

import "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"

type IOrderRepo interface {
	CreateOrder(order Entities.Order) (Entities.Order, error)
	ChangeOrderStatus(order Entities.Order, orderID string) (Entities.Order, error)
	GetOrderByID(orderID string) (Entities.Order, error)
}
