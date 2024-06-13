package Server

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Handlers"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
)

func (s *FiberServer) SetupRoute(uProduct UseCases.IProductCase, uStocks UseCases.IStockCase) {
	ProductHandler := Handlers.NewProductHandler(uProduct, uStocks)
	StockHandler := Handlers.NewStockHandler(uStocks, uProduct)

	s.app.Get("/products/", ProductHandler.GetAllProducts)
	s.app.Get("/product/:id", ProductHandler.GetProductById)
	s.app.Post("/createProduct/", ProductHandler.CreateProduct)
	s.app.Put("/updateProduct/:id", ProductHandler.UpdateProductById)
	s.app.Delete("/deleteProduct/:id", ProductHandler.DeleteProductById)

	s.app.Get("/stocks/", StockHandler.GetAllStock)
	s.app.Get("/stock/:id", StockHandler.GetStockByID)
	s.app.Post("/createStock/", StockHandler.CreateStock)
	s.app.Put("/updateStock/:id", StockHandler.UpdateStock)
	s.app.Delete("/deleteStock/:id", StockHandler.DeleteStock)
}
