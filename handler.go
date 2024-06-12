package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePerson(c *gin.Context) {

	var personReq PersonRequest
	if err := c.ShouldBindJSON(&personReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	_, err := db.Exec("INSERT INTO person(name, age) VALUES(?, ?)", personReq.Name, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create person"})
		return
	}

	var lastInsertID int
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&lastInsertID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get person ID"})
		return
	}

	_, err = db.Exec("INSERT INTO phone(person_id, number) VALUES(?, ?)", lastInsertID, personReq.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create phone record"})
		return
	}

	_, err = db.Exec("INSERT INTO address(city, state, street1, street2, zip_code) VALUES(?, ?, ?, ?, ?)",
		personReq.City, personReq.State, personReq.Street1, personReq.Street2, personReq.ZipCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create address record"})
		return
	}

	var lastAddressInsertID int
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&lastAddressInsertID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get address ID"})
		return
	}

	_, err = db.Exec("INSERT INTO address_join(person_id, address_id) VALUES(?, ?)", lastInsertID, lastAddressInsertID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create address_join record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person created successfully"})
}
func GetPersonInfo(c *gin.Context) {
	// Extract person_id from URL parameter
	personID, err := strconv.Atoi(c.Param("person_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid person ID"})
		return
	}

	query := `
        SELECT p.name, ph.number, a.city, a.state, a.street1, a.street2, a.zip_code
        FROM person p
        JOIN phone ph ON p.id = ph.person_id
        JOIN address_join aj ON p.id = aj.person_id
        JOIN address a ON aj.address_id = a.id
        WHERE p.id = ?
    `

	var (
		name     string
		number   string
		city     string
		state    string
		street1  string
		street2  string
		zip_code string
	)

	err = db.QueryRow(query, personID).Scan(&name, &number, &city, &state, &street1, &street2, &zip_code)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch person info"})
		}
		return
	}

	response := gin.H{
		"name":         name,
		"phone_number": number,
		"city":         city,
		"state":        state,
		"street1":      street1,
		"street2":      street2,
		"zip_code":     zip_code,
	}

	// Return response as JSON
	c.JSON(http.StatusOK, response)
}
