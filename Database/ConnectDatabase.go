package Database

import (
	"fmt"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

var db *gorm.DB

func InitDatabase() *gorm.DB {
	err := godotenv.Load(".env.example")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	host := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")

	dbPort, err := strconv.Atoi(databasePort)
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, dbPort, username, password, databaseName)
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected to database")
	db.AutoMigrate(&Model.Product{})
	db.AutoMigrate(&Model.Address{})
	db.AutoMigrate(&Model.Stock{})
	db.AutoMigrate(&Model.OrderCalculate{})
	db.AutoMigrate(&Model.Order{})
	return db
}
