package Server

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/MiddleWare"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/PKG/Handlers/Address"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/PKG/Handlers/Order"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/PKG/Handlers/Product"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/PKG/Handlers/Stock"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/PKG/Handlers/Transaction"
	UseCase "github.com/SHERATONS/OMS-Sellsuki-Internship/PKG/UseCases"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func (s *FiberServer) SetupRoute(uProduct UseCase.IProductUseCase, uStock UseCase.IStockUseCase, uAddress UseCase.IAddressUseCase, uTransactionID UseCase.ITransactionIDUseCase, uOrder UseCase.IOrderUseCase) {
	ProductHandler := Product.NewProductHandler(uProduct)
	StockHandler := Stock.NewStockHandler(uStock)
	AddressHandler := Address.NewAddressHandler(uAddress)
	TransactionIDHandler := Transaction.NewTransactionIDHandler(uTransactionID)
	OrderHandler := Order.NewOrderHandler(uOrder)

	s.app.Use(MiddleWare.TracingMiddleWare)
	s.app.Use(MiddleWare.LoggerMiddleWare)

	s.app.Get("/metricsx", func(c *fiber.Ctx) error {
		promHandler := promhttp.Handler()
		fasthttpadaptor.NewFastHTTPHandler(promHandler)(c.Context())
		return nil
	})

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
