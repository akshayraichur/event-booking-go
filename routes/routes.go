package routes

import (
	"fmt"
	"net/http"

	"akshayraichur.com/event-booking-go/utils"
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

	server.POST("/events", protectedRoutes(createEvent))

	server.PUT("/events/:id", protectedRoutes(updateEvent))

	server.DELETE("/events/:id", protectedRoutes(deleteEvent))

	// User routes
	server.POST("/signup", signUp)
	server.POST("/login", login)
	
}

func protectedRoutes(callbackFunction func(context *gin.Context)) gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")

		if token == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		err := utils.VerifyToken(token)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		callbackFunction(context)
	}
}