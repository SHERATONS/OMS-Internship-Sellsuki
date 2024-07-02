package Stock

import (
	"context"
	Stock2 "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Stock"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository/Product"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository/Stock"
	"time"
)

type StockUseCase struct {
	Repo        Stock.IStockRepo
	RepoProduct Product.IProductRepo
}

func (s StockUseCase) GetAllStocks(ctx context.Context) ([]Stock2.Stock, error) {
	ctx, span := tracer.Start(ctx, "GetAllStocks_UseCase")
	defer span.End()

	return s.Repo.GetAllStocks(ctx)
}

func (s StockUseCase) GetStockByID(ctx context.Context, stockID string) (Stock2.Stock, error) {
	ctx, span := tracer.Start(ctx, "GetStockByID_UseCase")
	defer span.End()
	return s.Repo.GetStockByID(ctx, stockID)
}

func (s StockUseCase) CreateStock(ctx context.Context, Stock Stock2.Stock) (Stock2.Stock, error) {
	ctx, span := tracer.Start(ctx, "CreateStock_UseCase")
	defer span.End()

	_, err := s.RepoProduct.GetProductByID(ctx, Stock.SID)
	if err != nil {
		return Stock, err
	}

	Stock.SUpdated = time.Now()

	return s.Repo.CreateStock(ctx, Stock)
}

func (s StockUseCase) UpdateStock(ctx context.Context, Stock Stock2.Stock, stockID string) (Stock2.Stock, error) {
	ctx, span := tracer.Start(ctx, "UpdateStock_UseCase")
	defer span.End()

	_, err := s.Repo.GetStockByID(ctx, stockID)
	if err != nil {
		return Stock, err
	}

	_, err = s.RepoProduct.GetProductByID(ctx, stockID)
	if err != nil {
		return Stock, err
	}

	Stock.SUpdated = time.Now()

	return s.Repo.UpdateStock(ctx, Stock, stockID)
}

func (s StockUseCase) DeleteStock(ctx context.Context, stockID string) error {
	ctx, span := tracer.Start(ctx, "DeleteStock_UseCase")
	defer span.End()

	_, err := s.Repo.GetStockByID(ctx, stockID)
	if err != nil {
		return err
	}

	return s.Repo.DeleteStock(ctx, stockID)
}

func NewStockUseCase(Repo Stock.IStockRepo, RepoProduct Product.IProductRepo) IStockUseCase {
	return StockUseCase{Repo: Repo, RepoProduct: RepoProduct}
}
