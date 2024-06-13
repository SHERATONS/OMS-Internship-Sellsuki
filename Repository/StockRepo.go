package Repository

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"gorm.io/gorm"
)

type StockRepo struct {
	Db *gorm.DB
}

func (s *StockRepo) GetAllStocks() ([]Entities.Stock, error) {
	var stocks []Entities.Stock
	err := s.Db.Order("s_id").Find(&stocks).Error
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

//func (s StockRepo) UpdateStock(Stock Entities.Stock) (Entities.Stock, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s StockRepo) DeleteStock(stockId string) (Entities.Stock, error) {
//	//TODO implement me
//	panic("implement me")
//}

func NewStockRepo(db *gorm.DB) IStockRepo {
	return &StockRepo{Db: db}
}
