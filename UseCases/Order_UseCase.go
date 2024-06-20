package UseCases

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository"
)

type OrderUseCase struct {
	Repo Repository.IOrderRepo
}

func (o OrderUseCase) GetOrderById(orderID string) (Entities.Order, error) {
	return o.Repo.GetOrderByID(orderID)
}

func (o OrderUseCase) CreateOrder(order Entities.Order) (Entities.Order, error) {
	return o.Repo.CreateOrder(order)
}

func (o OrderUseCase) ChangeOrderStatus(order Entities.Order, orderID string) (Entities.Order, error) {
	return o.Repo.ChangeOrderStatus(order, orderID)
}

func NewOrderUseCase(Repo Repository.IOrderRepo) IOrderUseCase {
	return &OrderUseCase{Repo: Repo}
}
