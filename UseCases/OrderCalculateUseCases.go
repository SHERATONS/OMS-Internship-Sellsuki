package UseCases

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository"
)

type OrderCalculateUseCases struct {
	Repo Repository.IOrderCalculateRepo
}

//func (o OrderCalculateUseCases) GetAllOrders() ([]Entities.OrderCalculate, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (o OrderCalculateUseCases) GetOrderByTransactionID(transactionID string) (Entities.OrderCalculate, error) {
//	//TODO implement me
//	panic("implement me")
//}

func (o OrderCalculateUseCases) CreateTransactionID(orderCalculate Entities.OrderCalculate) (Entities.OrderCalculate, error) {
	return o.Repo.CreateTransactionID(orderCalculate)
}

//func (o OrderCalculateUseCases) DeleteTransactionID(transactionID string) error {
//	//TODO implement me
//	panic("implement me")
//}

func NewOrderCalculateUseCases(repo Repository.IOrderCalculateRepo) IOrderCalculateCase {
	return OrderCalculateUseCases{Repo: repo}
}
