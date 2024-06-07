package UseCases

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
)

type OrderRepository interface {
	SaveOrder(order Entities.Order) error
}
