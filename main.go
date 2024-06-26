package main

import (
	"log"
	"os"

	"github.com/SHERATONS/OMS-Sellsuki-Internship/Database"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Server"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env.example")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	port := os.Getenv("PORT")

	db := Database.InitDatabase()

	// init product repo, use cases
	ProductRP := Repository.NewProductRepo(db)
	ProductUS := UseCases.NewProductUseCase(ProductRP)

	// init stock repo, use cases
	StockRP := Repository.NewStockRepo(db)
	StockUS := UseCases.NewStockUseCase(StockRP)

	// init address repo, use cases
	AddressRP := Repository.NewAddressRepo(db)
	AddressUs := UseCases.NewAddressUseCase(AddressRP)

	// init order calculation repo, use cases
	OrderCalculateRP := Repository.NewTransactionIDRepo(db)
	OrderCalculateUs := UseCases.NewTransactionIDUseCase(OrderCalculateRP)

	// init order repo, use cases
	OrderRP := Repository.NewOrderRepo(db)
	OrderUS := UseCases.NewOrderUseCase(OrderRP)

	s := Server.NewFiberServer()
	s.SetupRoute(ProductUS, StockUS, AddressUs, OrderCalculateUs, OrderUS)
	s.Start(port)
}
