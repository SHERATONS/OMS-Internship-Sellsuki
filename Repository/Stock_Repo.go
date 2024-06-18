package Repository

import (
	"errors"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Model"
	"gorm.io/gorm"
	"log"
)

type StockRepo struct {
	Db *gorm.DB
}

func (s *StockRepo) GetAllStocks() ([]Entities.Stock, error) {
	var stocks []Entities.Stock
	err := s.Db.Order("CAST(s_id AS INTEGER)").Find(&stocks).Error
	return stocks, err
}

func (s StockRepo) GetStockByID(stockId string) (Entities.Stock, error) {
	var stock Entities.Stock
	err := s.Db.Where("s_id = ?", stockId).First(&stock).Error
	return stock, err
}

func (s StockRepo) CreateStock(Stock Entities.Stock) (Entities.Stock, error) {
	err := s.Db.Create(&Stock).Error
	return Stock, err
}

func (s StockRepo) UpdateStock(Stock Entities.Stock, stockId string) (Entities.Stock, error) {
	if Stock.SQuantity < 0 {
		return Stock, errors.New("stock quantity is negative")
	}
	err := s.Db.Where("s_id = ?", stockId).Save(&Stock).Error
	return Stock, err
}

func (s StockRepo) DeleteStock(stockId string) error {
	err := s.Db.Where("s_id = ?", stockId).Delete(&Entities.Stock{}).Error
	return err
}

func NewStockRepo(db *gorm.DB) IStockRepo {
	err := db.AutoMigrate(&Model.Stock{})
	if err != nil {
		log.Fatalf("Failed to auto migrate Stock: %v", err)
	}
	return &StockRepo{Db: db}
}
