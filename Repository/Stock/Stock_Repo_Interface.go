package Stock

import (
	"context"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Stock"
	"go.opentelemetry.io/otel"
)

type IStockRepo interface {
	GetAllStocks(ctx context.Context) ([]Stock.Stock, error)
	GetStockByID(ctx context.Context, stockId string) (Stock.Stock, error)
	CreateStock(ctx context.Context, stock Stock.Stock) (Stock.Stock, error)
	UpdateStock(ctx context.Context, Stock Stock.Stock, stockID string) (Stock.Stock, error)
	DeleteStock(ctx context.Context, stockId string) error
}

var tracer = otel.Tracer("Stock_Repo")
