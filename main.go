package main

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Database"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/FiberServer"
	_ "github.com/lib/pq"
	//"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	//"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
)

func main() {
	Database.InitDatabase()
	FiberServer.ConnectFiberServer()
}
