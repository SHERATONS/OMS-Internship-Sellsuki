package Server

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Handlers"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
)

func (s *FiberServer) SetupRoute(uProduct UseCases.IProductCase, uStocks UseCases.IStockCase, uAddress UseCases.IAddressCase, uOrderCalculate UseCases.IOrderCalculateCase, uOrder UseCases.IOrderCase) {
	ProductHandler := Handlers.NewProductHandler(uProduct, uStocks)
	StockHandler := Handlers.NewStockHandler(uStocks, uProduct)
	AddressHandler := Handlers.NewAddressHandler(uAddress)
	OrderCalculateHandler := Handlers.NewOrderCalculateHandler(uOrderCalculate, uProduct, uAddress)
	OrderHandler := Handlers.NewOrderHandler(uOrder, uStocks)

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

	s.app.Get("/address/:city", AddressHandler.GetAddressByCity)
	s.app.Post("createAddress/", AddressHandler.CreateAddress)
	s.app.Put("updateAddress/:city", AddressHandler.UpdateAddress)
	s.app.Delete("deleteAddress/:city", AddressHandler.DeleteAddress)

	s.app.Get("/OrderCalculates/", OrderCalculateHandler.GetAllOrder)
	s.app.Get("/OrderCalculate/:tid", OrderCalculateHandler.GetOrderByTransactionID)
	s.app.Post("/order/calculate/", OrderCalculateHandler.CreateTransactionID)
	s.app.Delete("/deleteTransactionId/:tid", OrderCalculateHandler.DeleteTransactionID)

	s.app.Post("/order/", OrderHandler.CreateOrder)
	s.app.Patch("/order/status/:oid", OrderHandler.ChangeOrderStatus)
}
