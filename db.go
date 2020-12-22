package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB
var err error

func connectDatabase() {
	var envs map[string]string
	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := envs["DB_USER"]
	password := envs["DB_PASSWORD"]
	database := envs["DB_NAME"]

	db, err = sql.Open("mysql", user+":"+password+"@tcp(127.0.0.1:3306)/"+database+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established at http://localhost:8000/")
		log.Println("\nTo stop server: CONTROL + C")
	}
}
