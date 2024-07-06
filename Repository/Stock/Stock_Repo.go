package Stock

import (
	"context"
	"errors"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Stock"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Model"
	"gorm.io/gorm"
	"log"
	"time"
)

type StockRepo struct {
	Db *gorm.DB
}

func (s *StockRepo) GetAllStocks(ctx context.Context) ([]Stock.Stock, error) {
	ctx, span := tracer.Start(ctx, "GetAllStocks_Repo")
	defer span.End()

	var stocks []Stock.Stock

	err := s.Db.Order("CAST(s_id AS INTEGER)").Find(&stocks).Error
	if err != nil {
		return stocks, err
	}

	return stocks, nil
}

func (s StockRepo) GetStockByID(ctx context.Context, stockId string) (Stock.Stock, error) {
	ctx, span := tracer.Start(ctx, "GetStockByID_Repo")
	defer span.End()

	var stock Stock.Stock

	err := s.Db.Where("s_id = ?", stockId).First(&stock).Error
	if err != nil {
		return stock, errors.New("stock Not Found")
	}

	return stock, nil
}

func (s StockRepo) CreateStock(ctx context.Context, stock Stock.Stock) (Stock.Stock, error) {
	ctx, span := tracer.Start(ctx, "CreateStock_Repo")
	defer span.End()

	stock.SUpdated = time.Now()

	err := s.Db.Create(&stock).Error
	if err != nil {
		return stock, errors.New("failed to create stock")
	}

	return stock, nil
}

func (s StockRepo) UpdateStock(ctx context.Context, stock Stock.Stock, stockId string) (Stock.Stock, error) {
	ctx, span := tracer.Start(ctx, "UpdateStock_Repo")
	defer span.End()

	if stock.SQuantity < 0 {
		return stock, errors.New("stock quantity is negative")
	}

	stock.SUpdated = time.Now()

	err := s.Db.Where("s_id = ?", stockId).Save(&stock).Error
	if err != nil {
		return stock, errors.New("failed to update stock")
	}

	return stock, nil
}

func (s StockRepo) DeleteStock(ctx context.Context, stockId string) error {
	ctx, span := tracer.Start(ctx, "DeleteStock_Repo")
	defer span.End()

	err := s.Db.Where("s_id = ?", stockId).Delete(&Stock.Stock{}).Error
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
