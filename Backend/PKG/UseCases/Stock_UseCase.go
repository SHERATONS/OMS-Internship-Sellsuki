package UseCases

import (
	"context"
	Stock2 "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Stock"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/PKG/Repository/Product"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/PKG/Repository/Stock"
)

type StockUseCase struct {
	Repo        Stock.IStockRepo
	RepoProduct Product.IProductRepo
}

func (s StockUseCase) GetAllStocks(ctx context.Context) ([]Stock2.Stock, error) {
	ctx, span := tracerStock.Start(ctx, "GetAllStocks_UseCase")
	defer span.End()

	return s.Repo.GetAllStocks(ctx)
}

func (s StockUseCase) GetStockByID(ctx context.Context, stockID string) (Stock2.Stock, error) {
	ctx, span := tracerStock.Start(ctx, "GetStockByID_UseCase")
	defer span.End()

	return s.Repo.GetStockByID(ctx, stockID)
}

func (s StockUseCase) CreateStock(ctx context.Context, stock Stock2.Stock) (Stock2.Stock, error) {
	ctx, span := tracerStock.Start(ctx, "CreateStock_UseCase")
	defer span.End()

	_, err := s.RepoProduct.GetProductByID(ctx, stock.SID)
	if err != nil {
		return stock, err
	}

	return s.Repo.CreateStock(ctx, stock)
}

func (s StockUseCase) UpdateStock(ctx context.Context, stock Stock2.Stock, stockID string) (Stock2.Stock, error) {
	ctx, span := tracerStock.Start(ctx, "UpdateStock_UseCase")
	defer span.End()

	_, err := s.Repo.GetStockByID(ctx, stockID)
	if err != nil {
		return stock, err
	}

	_, err = s.RepoProduct.GetProductByID(ctx, stockID)
	if err != nil {
		return stock, err
	}

	return s.Repo.UpdateStock(ctx, stock, stockID)
}

func (s StockUseCase) DeleteStock(ctx context.Context, stockID string) error {
	ctx, span := tracerStock.Start(ctx, "DeleteStock_UseCase")
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
