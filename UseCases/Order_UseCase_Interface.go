package UseCases

import "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"

type IOrderUseCase interface {
	CreateOrder(order Entities.Order) (Entities.Order, error)
	ChangeOrderStatus(order Entities.Order, oid string) (Entities.Order, error)
	GetOrderById(orderId string) (Entities.Order, error)
}
