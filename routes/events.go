package routes

import (
	"github/rigel-developer/go-final-rest-api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mgutz/ansi"
)

var red = ansi.ColorFunc("red")

func getEvents(c *gin.Context) {

	events, err := models.GetAll()
	if err != nil {
		log.Println(red(err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting events"})
		return
	}

	c.JSON(200, gin.H{
		"status": http.StatusOK,
		"events": events,
	})
}

func getEvent(c *gin.Context) {
	// red := ansi.ColorFunc("red")
	id := c.Param("id")
	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println(red(err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	event, err := models.GetOne(newId)
	if err != nil {
		log.Println(red(err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting event"})
		return
	}

	c.JSON(200, gin.H{
		"status": http.StatusOK,
		"event":  event,
	})
}

func createEvent(context *gin.Context) {
	userIdToken := context.GetInt64("userId")
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		log.Println(red(err.Error()))
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	event.UserID = userIdToken

	err = event.Save()
	if err != nil {
		log.Println(red(err.Error()))
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Event created successfully!",
		"event":   event,
	})
}

func updateEvent(context *gin.Context) {
	userIdToken := context.GetInt64("userId")
	id := context.Param("id")
	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println(red(err.Error()))
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	e, err := models.GetOne(newId)
	if err != nil {
		log.Println(red(err.Error()))
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting event"})
		return
	}

	if e.UserID != userIdToken {
		context.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this event"})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		log.Println(red(err.Error()))
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	event.ID = newId
	err = event.Update()
	if err != nil {
		log.Println(red(err.Error()))
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Event updated successfully!",
		"event":   event,
	})
}

func deleteEvent(context *gin.Context) {
	userIdToken := context.GetInt64("userId")
	id := context.Param("id")
	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println(red(err.Error()))
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	event, err := models.GetOne(newId)
	if err != nil {
		log.Println(red(err.Error()))
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting event"})
		return
	}

	if event.UserID != userIdToken {
		context.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this event"})
		return
	}

	err = models.Delete(event.ID)
	if err != nil {
		log.Println(red(err.Error()))
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Event deleted successfully!",
	})
}
