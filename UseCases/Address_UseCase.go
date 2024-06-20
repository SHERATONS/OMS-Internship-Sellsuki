package UseCases

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository"
)

type AddressUseCase struct {
	Repo Repository.IAddressRepo
}

func (a AddressUseCase) GetAddressByCity(city string) (Entities.Address, error) {
	return a.Repo.GetAddressByCity(city)
}

func (a AddressUseCase) CreateAddress(address Entities.Address) (Entities.Address, error) {
	return a.Repo.CreateAddress(address)
}

func (a AddressUseCase) UpdateAddress(address Entities.Address, city string) (Entities.Address, error) {
	return a.Repo.UpdateAddress(address, city)
}

func (a AddressUseCase) DeleteAddress(city string) error {
	return a.Repo.DeleteAddress(city)
}

func NewAddressUseCase(repo Repository.IAddressRepo) IAddressUseCase {
	return &AddressUseCase{Repo: repo}
}
