package UseCases

import "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"

type OrderCalculation interface {
	OrderCalculation(order Entities.Order)
}
