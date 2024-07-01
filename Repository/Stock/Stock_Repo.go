package Stock

import (
	"context"
	"errors"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Model"
	"gorm.io/gorm"
	"log"
)

type StockRepo struct {
	Db *gorm.DB
}

func (s *StockRepo) GetAllStocks(ctx context.Context) ([]Entities.Stock, error) {
	//_, span := Stracer.Start(ctx, "GetAllStocks")
	//defer span.End()

	var stocks []Entities.Stock

	err := s.Db.Order("CAST(s_id AS INTEGER)").Find(&stocks).Error
	if err != nil {
		//span.RecordError(err)
		return stocks, err
	}

	return stocks, nil
}

func (s StockRepo) GetStockByID(ctx context.Context, stockId string) (Entities.Stock, error) {
	var stock Entities.Stock

	err := s.Db.Where("s_id = ?", stockId).First(&stock).Error
	if err != nil {
		return stock, errors.New("stock Not Found")
	}

	return stock, nil
}

func (s StockRepo) CreateStock(ctx context.Context, Stock Entities.Stock) (Entities.Stock, error) {
	err := s.Db.Create(&Stock).Error
	if err != nil {
		return Stock, errors.New("failed to create stock")
	}

	return Stock, nil
}

func (s StockRepo) UpdateStock(ctx context.Context, Stock Entities.Stock, stockId string) (Entities.Stock, error) {
	if Stock.SQuantity < 0 {
		return Stock, errors.New("stock quantity is negative")
	}

	err := s.Db.Where("s_id = ?", stockId).Save(&Stock).Error
	if err != nil {
		return Stock, errors.New("failed to update stock")
	}

	return Stock, nil
}

func (s StockRepo) DeleteStock(ctx context.Context, stockId string) error {
	err := s.Db.Where("s_id = ?", stockId).Delete(&Entities.Stock{}).Error
	if err != nil {
		return errors.New("failed to delete stock")
	}

	return err
}

func NewStockRepo(db *gorm.DB) IStockRepo {
	err := db.AutoMigrate(&Model.Stock{})
	if err != nil {
		log.Fatalf("Failed to auto migrate Stock: %v", err)
	}
	return &StockRepo{Db: db}
}
