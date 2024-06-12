package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Setup database connection
	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Define routes
	router.GET("/person/:id", func(c *gin.Context) {
		getPersonDetails(c, db)
	})

	// Start the server
	router.Run(":8080")
}
