package main

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Observability"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository/Address"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository/Order"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository/Product"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository/Stock"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository/Transaction"
	Address2 "github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases/Address"
	Order2 "github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases/Order"
	Product2 "github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases/Product"
	Stock2 "github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases/Stock"
	Transaction2 "github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases/Transaction"
	"log"
	"os"

	"github.com/SHERATONS/OMS-Sellsuki-Internship/Database"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := Observability.InitTracer(); err != nil {
		log.Fatal(err)
	}

	err := godotenv.Load(".env.example")
	if err != nil {
		log.Fatal("Error loading .env file", err)
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
	ProductUS := Product2.NewProductUseCase(ProductRP, StockRP)
	StockUS := Stock2.NewStockUseCase(StockRP, ProductRP)
	AddressUs := Address2.NewAddressUseCase(AddressRP)
	TransactionIDUs := Transaction2.NewTransactionIDUseCase(TransactionIDRP, ProductRP, AddressRP)
	OrderUS := Order2.NewOrderUseCase(OrderRP, StockRP, TransactionIDRP)

	s := Server.NewFiberServer()
	s.SetupRoute(ProductUS, StockUS, AddressUs, TransactionIDUs, OrderUS)
	s.Start(port)
}
