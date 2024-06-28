package main

import (
	"github/rigel-developer/go-final-rest-api/db"
	"github/rigel-developer/go-final-rest-api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mgutz/ansi"
)

var red = ansi.ColorFunc("red")

func main() {

	// Initialize the database
	db.InitDatabase()

	// Create a new instance of the server
	server := gin.Default()

	//colors

	// green := ansi.ColorFunc("green")

	// yellow := ansi.ColorFunc("yellow")

	// //middleware
	// server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	statusColor := green

	// 	if param.StatusCode >= 400 {
	// 		statusColor = red
	// 	} else if param.StatusCode >= 300 {
	// 		statusColor = yellow
	// 	}

	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s\" %s %s \"%s\" %s\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC1123),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		statusColor(strconv.Itoa(param.StatusCode)),
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.ErrorMessage,
	// 	)
	// }))

	server.Use(gin.Recovery())

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting event"})
		return
	}

	c.JSON(200, gin.H{
		"status": http.StatusOK,
		"event":  event,
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
