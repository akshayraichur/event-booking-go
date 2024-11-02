package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"akshayraichur.com/event-booking-go/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(context *gin.Context) {

	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Could not fetch events"})
		return
	}

	context.JSON(http.StatusOK, events)

}

func CreateEvent(context *gin.Context) {

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

func GetEvent(context *gin.Context) {
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
