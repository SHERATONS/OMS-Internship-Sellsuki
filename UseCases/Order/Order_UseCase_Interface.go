package Order

import "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"

type IOrderUseCase interface {
	CreateOrder(TransactionID string) (Entities.Order, error)
	ChangeOrderStatus(oid string, oStatus string) (Entities.Order, error)
	GetOrderById(orderId string) (Entities.Order, error)
}
