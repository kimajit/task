package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func setupDB() (*sql.DB, error) {
	dsn := "username:password@tcp(127.0.0.1:3306)/cetec"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connection established")
	return db, nil
}
