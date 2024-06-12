package main

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Database"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Server"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
	_ "github.com/lib/pq"
	//"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	//"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
)

func main() {
	db := Database.InitDatabase()
	ProductRP := Repository.NewProductRepo(db)
	ProductUS := UseCases.NewProductUseCases(ProductRP)
	s := Server.NewFiberServer()
	s.SetupRoute(ProductUS)
	s.Start(":8080")
}
