package Server

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Handlers"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
)

func (s *FiberServer) SetupRoute(uProduct UseCases.IProductUseCase, uStock UseCases.IStockUseCase, uAddress UseCases.IAddressUseCase, uTransactionID UseCases.ITransactionIDUseCase, uOrder UseCases.IOrderUseCase) {
	ProductHandler := Handlers.NewProductHandler(uProduct, uStock)
	StockHandler := Handlers.NewStockHandler(uStock, uProduct)
	AddressHandler := Handlers.NewAddressHandler(uAddress)
	TransactionIDHandler := Handlers.NewTransactionIDHandler(uTransactionID, uProduct, uAddress)
	OrderHandler := Handlers.NewOrderHandler(uOrder, uStock, uTransactionID)

	// Product Route
	s.app.Get("/products/", ProductHandler.GetAllProducts)
	s.app.Get("/product/:id", ProductHandler.GetProductByID)
	s.app.Post("/product/create/", ProductHandler.CreateProduct)
	s.app.Put("/product/update/:id", ProductHandler.UpdateProductById)
	s.app.Delete("/product/delete/:id", ProductHandler.DeleteProductById)

	// Stock Route
	s.app.Get("/stocks/", StockHandler.GetAllStock)
	s.app.Get("/stock/:id", StockHandler.GetStockByID)
	s.app.Post("/stock/create/", StockHandler.CreateStock)
	s.app.Put("/stock/update/:id", StockHandler.UpdateStock)
	s.app.Delete("/stock/delete/:id", StockHandler.DeleteStock)

	// Address Route
	s.app.Get("/address/:city", AddressHandler.GetAddressByCity)
	s.app.Post("/address/create/", AddressHandler.CreateAddress)
	s.app.Put("/address/update/:city", AddressHandler.UpdateAddress)
	s.app.Delete("address/delete/:city", AddressHandler.DeleteAddress)

	// TransactionID Route
	s.app.Get("/transactionIDs/", TransactionIDHandler.GetAllTransactionIDs)
	s.app.Get("/transactionID/:tid", TransactionIDHandler.GetOrderByTransactionID)
	s.app.Post("/order/calculate/", TransactionIDHandler.CreateTransactionID)
	s.app.Delete("/transactionID/delete/:tid", TransactionIDHandler.DeleteTransactionID)

	// Order Route
	s.app.Get("/order/:oid", OrderHandler.GetOrderById)
	s.app.Post("/order/create/", OrderHandler.CreateOrder)
	s.app.Patch("/order/status/:oid", OrderHandler.ChangeOrderStatus)
}
