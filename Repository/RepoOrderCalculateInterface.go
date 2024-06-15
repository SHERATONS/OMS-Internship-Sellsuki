package Repository

import "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"

type IOrderCalculateRepo interface {
	//GetAllOrders() ([]Entities.OrderCalculate, error)
	//GetOrderByTransactionID(transactionID string) (Entities.OrderCalculate, error)
	CreateTransactionID(orderCalculate Entities.OrderCalculate) (Entities.OrderCalculate, error)
	//DeleteTransactionID(transactionID string) error
}
