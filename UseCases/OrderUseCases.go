package UseCases

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository"
)

type OrderUseCases struct {
	Repo Repository.IOrderRepo
}

func (o OrderUseCases) GetOrderById(orderId int64) (Entities.Order, error) {
	return o.Repo.GetOrderById(orderId)
}

func (o OrderUseCases) CreateOrder(order Entities.Order) (Entities.Order, error) {
	return o.Repo.CreateOrder(order)
}

func (o OrderUseCases) ChangeOrderStatus(order Entities.Order, oid int64) (Entities.Order, error) {
	return o.Repo.ChangeOrderStatus(order, oid)
}

func NewOrderUseCases(Repo Repository.IOrderRepo) IOrderCase {
	return &OrderUseCases{Repo: Repo}
}
