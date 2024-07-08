package Address

import (
	"context"
	"errors"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Address"
	"log"
	"time"

	"github.com/SHERATONS/OMS-Sellsuki-Internship/Model"
	"gorm.io/gorm"
)

type AddressRepo struct {
	Db *gorm.DB
}

func (a AddressRepo) GetAddressByCity(ctx context.Context, city string) (Address.Address, error) {
	ctx, span := tracer.Start(ctx, "GetAddressByCity_Repo")
	defer span.End()

	var address Address.Address
	err := a.Db.Where("city = ?", city).First(&address).Error
	if err != nil {
		return address, errors.New("address not found")
	}

	return address, nil
}

func (a AddressRepo) CreateAddress(ctx context.Context, address Address.Address) (Address.Address, error) {
	ctx, span := tracer.Start(ctx, "CreateAddress_Repo")
	defer span.End()

	address.AUpdated = time.Now()

	err := a.Db.Create(&address).Error
	if err != nil {
		return address, errors.New("failed to create address")
	}

	return address, nil
}

func (a AddressRepo) UpdateAddress(ctx context.Context, address Address.Address, city string) (Address.Address, error) {
	ctx, span := tracer.Start(ctx, "UpdateAddress_Repo")
	defer span.End()

	address.AUpdated = time.Now()

	err := a.Db.Where("city = ?", city).Save(&address).Error
	if err != nil {
		return address, errors.New("failed to update address")
	}

	return address, nil
}

func (a AddressRepo) DeleteAddress(ctx context.Context, city string) error {
	ctx, span := tracer.Start(ctx, "DeleteAddress_Repo")
	defer span.End()

	err := a.Db.Where("city = ?", city).Delete(&Address.Address{}).Error
	if err != nil {
		return errors.New("failed to delete address")
	}

	return err
}

func NewAddressRepo(db *gorm.DB) IAddressRepo {
	err := db.AutoMigrate(&Model.Address{})
	if err != nil {
		log.Fatalf("Failed to auto migrate Address: %v", err)
	}
	return &AddressRepo{Db: db}
}
