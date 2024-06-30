package UseCases

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository"
	"time"
)

type StockUseCase struct {
	Repo        Repository.IStockRepo
	RepoProduct Repository.IProductRepo
}

func (s StockUseCase) GetAllStocks() ([]Entities.Stock, error) {
	return s.Repo.GetAllStocks()
}

func (s StockUseCase) GetStockByID(stockID string) (Entities.Stock, error) {
	return s.Repo.GetStockByID(stockID)
}

func (s StockUseCase) CreateStock(Stock Entities.Stock) (Entities.Stock, error) {
	_, err := s.RepoProduct.GetProductByID(Stock.SID)
	if err != nil {
		return Stock, err
	}

	Stock.SUpdated = time.Now()
	
	return s.Repo.CreateStock(Stock)
}

func (s StockUseCase) UpdateStock(Stock Entities.Stock, stockID string) (Entities.Stock, error) {
	_, err := s.Repo.GetStockByID(stockID)
	if err != nil {
		return Stock, err
	}

	_, err = s.RepoProduct.GetProductByID(stockID)
	if err != nil {
		return Stock, err
	}

	Stock.SUpdated = time.Now()

	return s.Repo.UpdateStock(Stock, stockID)
}

func (s StockUseCase) DeleteStock(stockID string) error {
	_, err := s.Repo.GetStockByID(stockID)
	if err != nil {
		return err
	}

	return s.Repo.DeleteStock(stockID)
}

func NewStockUseCase(Repo Repository.IStockRepo, RepoProduct Repository.IProductRepo) IStockUseCase {
	return StockUseCase{Repo: Repo, RepoProduct: RepoProduct}
}
