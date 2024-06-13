package UseCases

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository"
)

type StockUseCases struct {
	Repo Repository.IStockRepo
}

func (s StockUseCases) GetAllStocks() ([]Entities.Stock, error) {
	return s.Repo.GetAllStocks()
}

func (s StockUseCases) GetStockByID(stockId string) (Entities.Stock, error) {
	return s.Repo.GetStockByID(stockId)
}

func (s StockUseCases) CreateStock(Stock Entities.Stock) (Entities.Stock, error) {
	return s.Repo.CreateStock(Stock)
}

func (s StockUseCases) UpdateStock(Stock Entities.Stock, stockId string) (Entities.Stock, error) {
	return s.Repo.UpdateStock(Stock, stockId)
}

func (s StockUseCases) DeleteStock(stockId string) error {
	return s.Repo.DeleteStock(stockId)
}

func NewStockUseCases(Repo Repository.IStockRepo) IStockCase {
	return StockUseCases{Repo: Repo}
}
