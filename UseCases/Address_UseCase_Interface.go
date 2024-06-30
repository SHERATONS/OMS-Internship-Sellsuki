package UseCases

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
)

type IAddressUseCase interface {
	GetAddressByCity(city string) (Entities.Address, error)
	CreateAddress(address Entities.Address) (Entities.Address, error)
	UpdateAddress(address Entities.Address, city string) (Entities.Address, error)
	DeleteAddress(city string) error
}
