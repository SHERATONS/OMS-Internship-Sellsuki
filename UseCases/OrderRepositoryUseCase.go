package UseCases

import "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"

type OrderRepositoryUseCase interface {
	CreateOrder(order Entities.Order) error
}

type OrderService struct {
	repo OrderRepository
}

func NewOrder(repo OrderRepository) OrderRepositoryUseCase {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order Entities.Order) error {
	return s.repo.SaveOrder(order)
}
