package routes

import (
	"github/rigel-developer/go-final-rest-api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {

	userIdToken := c.GetInt64("userId")
	eventId := c.Param("id")
	newEventId, err := strconv.ParseInt(eventId, 10, 64)
	if err != nil {
		log.Println(red(err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Check if the event exists
	event, err := models.GetOne(newEventId)
	if err != nil {
		log.Println(red(err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting event"})
		return
	}

	var register models.Register
	register.EventID = event.ID
	register.UserID = userIdToken

	err = register.RegisterForEvent()
	if err != nil {
		log.Println(red(err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering for event"})
		return
	}

	c.JSON(200, gin.H{
		"status": http.StatusOK,
		"event":  event,
	})

}

func unregisterForEvent(c *gin.Context) {
	userIdToken := c.GetInt64("userId")
	eventId := c.Param("id")
	newEventId, err := strconv.ParseInt(eventId, 10, 64)
	if err != nil {
		log.Println(red(err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Check if the event exists
	event, err := models.GetOne(newEventId)
	if err != nil {
		log.Println(red(err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting event"})
		return
	}

	var register models.Register
	register.EventID = event.ID
	register.UserID = userIdToken

	err = models.CancelRegistration(userIdToken, event.ID)
	if err != nil {
		log.Println(red(err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error unregistering for event"})
		return
	}

	c.JSON(200, gin.H{
		"status": http.StatusOK,
		"event":  event,
	})

}

func getRegistrations(c *gin.Context) {
	userIdToken := c.GetInt64("userId")
	events, err := models.GetRegistrations(userIdToken)
	if err != nil {
		log.Println(red(err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting registrations"})
		return
	}

	c.JSON(200, gin.H{
		"status":        http.StatusOK,
		"registrations": events,
	})
}
