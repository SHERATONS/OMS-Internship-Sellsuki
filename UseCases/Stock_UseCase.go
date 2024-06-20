package UseCases

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository"
)

type StockUseCase struct {
	Repo Repository.IStockRepo
}

func (s StockUseCase) GetAllStocks() ([]Entities.Stock, error) {
	return s.Repo.GetAllStocks()
}

func (s StockUseCase) GetStockByID(stockID string) (Entities.Stock, error) {
	return s.Repo.GetStockByID(stockID)
}

func (s StockUseCase) CreateStock(Stock Entities.Stock) (Entities.Stock, error) {
	return s.Repo.CreateStock(Stock)
}

func (s StockUseCase) UpdateStock(Stock Entities.Stock, stockID string) (Entities.Stock, error) {
	return s.Repo.UpdateStock(Stock, stockID)
}

func (s StockUseCase) DeleteStock(stockID string) error {
	return s.Repo.DeleteStock(stockID)
}

func NewStockUseCase(Repo Repository.IStockRepo) IStockUseCase {
	return StockUseCase{Repo: Repo}
}
