package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const port = ":8080"

const (
	host         = "localhost"
	databasePort = 5432
	databaseName = "Order_week1"
	username     = "postgres"
	password     = "Ming1234"
)

var db *sql.DB

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, databasePort, username, password, databaseName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	print("database connected")
	//defer db.Close()
}
