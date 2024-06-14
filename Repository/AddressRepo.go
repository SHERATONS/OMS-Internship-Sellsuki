package Repository

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"gorm.io/gorm"
)

type AddressRepo struct {
	Db *gorm.DB
}

func (a AddressRepo) GetAddressByCity(city string) (Entities.Address, error) {
	var address Entities.Address
	err := a.Db.Where("city = ?", city).First(&address).Error
	return address, err
}

func (a AddressRepo) CreateAddress(address Entities.Address) (Entities.Address, error) {
	err := a.Db.Create(&address).Error
	return address, err
}

func (a AddressRepo) UpdateAddress(address Entities.Address, city string) (Entities.Address, error) {
	err := a.Db.Where("city = ?", city).Save(&address).Error
	return address, err
}

func (a AddressRepo) DeleteAddress(city string) error {
	err := a.Db.Where("city = ?", city).Delete(&Entities.Address{}).Error
	return err
}

func NewAddressRepo(db *gorm.DB) IAddressRepo {
	return &AddressRepo{Db: db}
}
