package Stock

import (
	"context"
	"errors"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Stock"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Model"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"log"
	"reflect"
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

	s.SetStockSubAttributes(stocks, span)

	return stocks, nil
}

func (s *StockRepo) GetStockByID(ctx context.Context, stockId string) (Stock.Stock, error) {
	ctx, span := tracer.Start(ctx, "GetStockByID_Repo")
	defer span.End()

	var stock Stock.Stock

	err := s.Db.Where("s_id = ?", stockId).First(&stock).Error
	if err != nil {
		return stock, errors.New("stock Not Found")
	}

	s.SetStockSubAttributes(stock, span)

	return stock, nil
}

func (s *StockRepo) CreateStock(ctx context.Context, stock Stock.Stock) (Stock.Stock, error) {
	ctx, span := tracer.Start(ctx, "CreateStock_Repo")
	defer span.End()

	stock.SUpdated = time.Now()

	err := s.Db.Create(&stock).Error
	if err != nil {
		return stock, errors.New("failed to create stock")
	}

	s.SetStockSubAttributes(stock, span)

	return stock, nil
}

func (s *StockRepo) UpdateStock(ctx context.Context, stock Stock.Stock, stockId string) (Stock.Stock, error) {
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

	s.SetStockSubAttributes(stock, span)

	return stock, nil
}

func (s *StockRepo) DeleteStock(ctx context.Context, stockId string) error {
	ctx, span := tracer.Start(ctx, "DeleteStock_Repo")
	defer span.End()

	err := s.Db.Where("s_id = ?", stockId).Delete(&Stock.Stock{}).Error
	if err != nil {
		return errors.New("failed to delete stock")
	}

	return err
}

func (s *StockRepo) SetStockSubAttributes(stockData any, sp trace.Span) {
	if stock, ok := stockData.(Stock.Stock); ok {
		sp.SetAttributes(
			attribute.String("StockID", stock.SID),
			attribute.Float64("StockQuantity", stock.SQuantity),
		)
	} else if stocks, ok := stockData.(*[]Stock.Stock); ok {
		stockIDs := make([]string, len(*stocks))
		stockQuantities := make([]float64, len(*stocks))

		for _, stock := range *stocks {
			stockIDs = append(stockIDs, stock.SID)
			stockQuantities = append(stockQuantities, stock.SQuantity)
		}

		sp.SetAttributes(
			attribute.StringSlice("StockID", stockIDs),
			attribute.Float64Slice("StockQuantity", stockQuantities),
		)
	} else {
		sp.RecordError(errors.New("invalid type: " + reflect.TypeOf(stockData).String()))
	}
}

func NewStockRepo(db *gorm.DB) IStockRepo {
	err := db.AutoMigrate(&Model.Stock{})
	if err != nil {
		log.Fatalf("Failed to auto migrate Stock: %v", err)
	}
	return &StockRepo{Db: db}
}
