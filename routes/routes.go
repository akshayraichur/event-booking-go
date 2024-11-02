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

	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)

	server.POST("/events", CreateEvent)
	
}