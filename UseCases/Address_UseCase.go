package UseCases

import (
	"context"
	Address2 "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Address"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository/Address"
)

type AddressUseCase struct {
	Repo Address.IAddressRepo
}

func (a AddressUseCase) GetAddressByCity(ctx context.Context, city string) (Address2.Address, error) {
	ctx, span := tracerAddress.Start(ctx, "GetAddressByCity_UseCase")
	defer span.End()

	return a.Repo.GetAddressByCity(ctx, city)
}

func (a AddressUseCase) CreateAddress(ctx context.Context, address Address2.Address) (Address2.Address, error) {
	ctx, span := tracerAddress.Start(ctx, "CreateAddress_UseCase")
	defer span.End()

	return a.Repo.CreateAddress(ctx, address)
}

func (a AddressUseCase) UpdateAddress(ctx context.Context, address Address2.Address, city string) (Address2.Address, error) {
	ctx, span := tracerAddress.Start(ctx, "UpdateAddress_UseCase")
	defer span.End()

	_, err := a.Repo.GetAddressByCity(ctx, city)
	if err != nil {
		return address, err
	}

	return a.Repo.UpdateAddress(ctx, address, city)
}

func (a AddressUseCase) DeleteAddress(ctx context.Context, city string) error {
	ctx, span := tracerAddress.Start(ctx, "DeleteAddress_UseCase")
	defer span.End()

	_, err := a.Repo.GetAddressByCity(ctx, city)
	if err != nil {
		return err
	}

	return a.Repo.DeleteAddress(ctx, city)
}

func NewAddressUseCase(repo Address.IAddressRepo) IAddressUseCase {
	return &AddressUseCase{Repo: repo}
}
