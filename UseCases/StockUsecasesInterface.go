package UseCases

import "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"

type IStockCase interface {
	GetAllStocks() ([]Entities.Stock, error)
	GetStockByID(stockId string) (Entities.Stock, error)
	CreateStock(Stock Entities.Stock) (Entities.Stock, error)
	//UpdateStock(Stock Entities.Stock) (Entities.Stock, error)
	//DeleteStock(stockId string) (Entities.Stock, error)
}
