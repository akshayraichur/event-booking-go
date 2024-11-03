package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	fmt.Println("Registering routes")
	
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	server.POST("/events", createEvent)

	server.PUT("/events/:id", updateEvent)

	server.DELETE("/events/:id", deleteEvent)
	
}