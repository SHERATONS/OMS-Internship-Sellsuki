package Address

import (
	"context"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"go.opentelemetry.io/otel"
)

type IAddressRepo interface {
	GetAddressByCity(ctx context.Context, city string) (Entities.Address, error)
	CreateAddress(ctx context.Context, address Entities.Address) (Entities.Address, error)
	UpdateAddress(ctx context.Context, address Entities.Address, city string) (Entities.Address, error)
	DeleteAddress(ctx context.Context, city string) error
}

var tracer = otel.Tracer("Address_Repo")
