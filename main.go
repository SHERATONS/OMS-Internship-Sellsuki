package main

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Database"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Server"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
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
	ProductUS := UseCases.NewProductUseCases(ProductRP)

	// init stock repo, use cases
	StockRP := Repository.NewStockRepo(db)
	StockUS := UseCases.NewStockUseCases(StockRP)

	// init address repo, use cases
	AddressRP := Repository.NewAddressRepo(db)
	AddressUs := UseCases.NewAddressUseCase(AddressRP)

	s := Server.NewFiberServer()
	s.SetupRoute(ProductUS, StockUS, AddressUs)
	s.Start(port)
}
