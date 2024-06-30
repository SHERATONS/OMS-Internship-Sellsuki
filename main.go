package main

import (
	"log"
	"os"

	"github.com/SHERATONS/OMS-Sellsuki-Internship/Database"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Server"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Trace"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	shutdown := Trace.InitTracer()
	defer shutdown()

	err := godotenv.Load(".env.example")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	port := os.Getenv("PORT")

	db := Database.InitDatabase()

	// init repo
	ProductRP := Repository.NewProductRepo(db)
	StockRP := Repository.NewStockRepo(db)
	AddressRP := Repository.NewAddressRepo(db)
	TransactionIDRP := Repository.NewTransactionIDRepo(db)
	OrderRP := Repository.NewOrderRepo(db)

	// init use case
	ProductUS := UseCases.NewProductUseCase(ProductRP, StockRP)
	StockUS := UseCases.NewStockUseCase(StockRP, ProductRP)
	AddressUs := UseCases.NewAddressUseCase(AddressRP)
	TransactionIDUs := UseCases.NewTransactionIDUseCase(TransactionIDRP, ProductRP, AddressRP)
	OrderUS := UseCases.NewOrderUseCase(OrderRP, StockRP, TransactionIDRP)

	s := Server.NewFiberServer()
	s.SetupRoute(ProductUS, StockUS, AddressUs, TransactionIDUs, OrderUS)
	s.Start(port)
}
