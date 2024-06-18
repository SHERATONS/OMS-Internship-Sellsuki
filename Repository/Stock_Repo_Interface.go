package Repository

import "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"

type IStockRepo interface {
	GetAllStocks() ([]Entities.Stock, error)
	GetStockByID(stockId string) (Entities.Stock, error)
	CreateStock(stock Entities.Stock) (Entities.Stock, error)
	UpdateStock(Stock Entities.Stock, stockID string) (Entities.Stock, error)
	DeleteStock(stockId string) error
}
