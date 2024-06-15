package UseCases

import "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"

type IOrderCalculateCase interface {
	//GetAllOrders() ([]Entities.OrderCalculate, error)
	//GetOrderByTransactionID(transactionID string) (Entities.OrderCalculate, error)
	CreateTransactionID(orderCalculate Entities.OrderCalculate) (Entities.OrderCalculate, error)
	//DeleteTransactionID(transactionID string) error
}
