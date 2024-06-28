package main

import (
	"github/rigel-developer/go-final-rest-api/db"
	"github/rigel-developer/go-final-rest-api/routes"

	"github.com/gin-gonic/gin"
)

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

	routes.RegisterRoutes(server)

	server.Run(":8080")

}
