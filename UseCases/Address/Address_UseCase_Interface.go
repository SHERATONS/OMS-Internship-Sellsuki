package Address

import (
	"context"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Address"
	"go.opentelemetry.io/otel"
)

type IAddressUseCase interface {
	GetAddressByCity(ctx context.Context, city string) (Address.Address, error)
	CreateAddress(ctx context.Context, address Address.Address) (Address.Address, error)
	UpdateAddress(ctx context.Context, address Address.Address, city string) (Address.Address, error)
	DeleteAddress(ctx context.Context, city string) error
}

var tracer = otel.Tracer("Address_UseCase")
