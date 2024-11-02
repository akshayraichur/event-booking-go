package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default() // sets up an HTTP server with some default middleware (with logger and recovery middleware)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/events", getEvents)

	server.Run(":8080") // listen and serve on

}

func getEvents(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"events": []string{"event1", "event2", "event3"},
	})

}
