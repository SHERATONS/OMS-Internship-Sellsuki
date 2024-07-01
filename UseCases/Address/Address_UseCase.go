package Address

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository/Address"
	"time"
)

type AddressUseCase struct {
	Repo Address.IAddressRepo
}

func (a AddressUseCase) GetAddressByCity(city string) (Entities.Address, error) {
	return a.Repo.GetAddressByCity(city)
}

func (a AddressUseCase) CreateAddress(address Entities.Address) (Entities.Address, error) {
	address.AUpdated = time.Now()

	return a.Repo.CreateAddress(address)
}

func (a AddressUseCase) UpdateAddress(address Entities.Address, city string) (Entities.Address, error) {
	_, err := a.Repo.GetAddressByCity(city)
	if err != nil {
		return address, err
	}

	address.AUpdated = time.Now()

	return a.Repo.UpdateAddress(address, city)
}

func (a AddressUseCase) DeleteAddress(city string) error {
	_, err := a.Repo.GetAddressByCity(city)
	if err != nil {
		return err
	}

	return a.Repo.DeleteAddress(city)
}

func NewAddressUseCase(repo Address.IAddressRepo) IAddressUseCase {
	return &AddressUseCase{Repo: repo}
}
