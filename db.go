package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
}
