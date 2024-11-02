package main

import (
	"fmt"
	"net/http"
	"strconv"

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
	server.GET("/events/:id", getEvent)

	server.POST("/events", createEvent)
	
	server.Run(":8080") // listen and serve on

}

func getEvents(context *gin.Context) {

	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Could not fetch events"})
		return
	}

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
	err = event.Save()
	
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Could not save event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": "Event created successfully", "event": event})

}


func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}
	
	event, err := models.GetEventById(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Could not fetch an event"})
		return
	}
	context.JSON(http.StatusOK, event)

}