package main

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Database"
	logger "github.com/SHERATONS/OMS-Sellsuki-Internship/Observability/Log"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Observability/Trace"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/PKG/Repository/Address"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/PKG/Repository/Order"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/PKG/Repository/Product"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/PKG/Repository/Stock"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/PKG/Repository/Transaction"
	UseCase "github.com/SHERATONS/OMS-Sellsuki-Internship/PKG/UseCases"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	logger.InitLogger()
	defer func(LogFile *os.File) {
		err := LogFile.Close()
		if err != nil {
		}
	}(logger.LogFile)

	fields := logrus.Fields{"module": "main", "function": "main"}
	logger.LogInfo("Service started", fields)

	err := Trace.InitTracer()
	if err != nil {
		logger.LogError("Failed to initialize trace"+err.Error(), fields)
	}

	if err := godotenv.Load(); err != nil {
	}

	port := os.Getenv("PORT")

	db := Database.InitDatabase()

	// init repo
	ProductRP := Product.NewProductRepo(db)
	StockRP := Stock.NewStockRepo(db)
	AddressRP := Address.NewAddressRepo(db)
	TransactionIDRP := Transaction.NewTransactionIDRepo(db)
	OrderRP := Order.NewOrderRepo(db)

	// init use case
	ProductUS := UseCase.NewProductUseCase(ProductRP, StockRP)
	StockUS := UseCase.NewStockUseCase(StockRP, ProductRP)
	AddressUs := UseCase.NewAddressUseCase(AddressRP)
	TransactionIDUs := UseCase.NewTransactionIDUseCase(TransactionIDRP, ProductRP, AddressRP)
	OrderUS := UseCase.NewOrderUseCase(OrderRP, StockRP, TransactionIDRP)

	s := Server.NewFiberServer()
	s.SetupRoute(ProductUS, StockUS, AddressUs, TransactionIDUs, OrderUS)
	err = s.Start(port)
	if err != nil {
		return
	}

}
