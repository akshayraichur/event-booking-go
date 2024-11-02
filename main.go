package main

import (
	"fmt"
	"net/http"

	"akshayraichur.com/event-booking-go/db"
	"akshayraichur.com/event-booking-go/models"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default() // sets up an HTTP server with some default middleware (with logger and recovery middleware)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") // listen and serve on

}

func getEvents(context *gin.Context) {

	events := models.GetAllEvents()

	context.JSON(http.StatusOK, events)

}

func createEvent(context *gin.Context) {

	var event models.Event
	fmt.Println("Checking this")

	// ShouldBindJSON is a helper function in gin that binds the request body to a struct

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()

	context.JSON(http.StatusCreated, gin.H{"status": "Event created successfully", "event": event})

}