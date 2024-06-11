package FiberServer

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"gorm.io/gorm"
	"log"
)

func GetProducts(db *gorm.DB) []Entities.Product {
	var product []Entities.Product
	result := db.Find(&product)

	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return product
}
