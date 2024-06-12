package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/person/:person_id/info", GetPersonInfo)
	router.POST("/person/create", CreatePerson)

	// Run server
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
