package main

import (
	"github/rigel-developer/go-final-rest-api/db"
	"github/rigel-developer/go-final-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize the database
	db.InitDatabase()

	// Create a new instance of the server
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")

}

func getEvents(c *gin.Context) {

	events, err := models.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting events"})
		return
	}

	c.JSON(200, gin.H{
		"status": http.StatusOK,
		"events": events,
	})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Event created successfully!",
		"event":   event,
	})
}
