package Stock

import (
	"context"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
)

type IStockRepo interface {
	GetAllStocks(ctx context.Context) ([]Entities.Stock, error)
	GetStockByID(ctx context.Context, stockId string) (Entities.Stock, error)
	CreateStock(ctx context.Context, stock Entities.Stock) (Entities.Stock, error)
	UpdateStock(ctx context.Context, Stock Entities.Stock, stockID string) (Entities.Stock, error)
	DeleteStock(ctx context.Context, stockId string) error
}
