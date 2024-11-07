package routes

import (
	"net/http"
	"strconv"

	"akshayraichur.com/event-booking-go/models"
	"github.com/gin-gonic/gin"
)

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

	// ShouldBindJSON is a helper function in gin that binds the request body to a struct

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := context.GetInt64("userID")

	event.UserID = userID

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

func updateEvent(context *gin.Context) {
	eventId := context.Param("id")
	
	eventID, err := strconv.ParseInt(eventId, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch an event"})
		return
	}
	userID := context.GetInt64("userID")

	if event.UserID != userID {
		context.JSON(http.StatusBadRequest, gin.H{ "message": "Not authorized to update"})
		return;
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "could not parse request body"})
		return
	}

	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Could not update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "Event updated successfully", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {

	eventId := context.Param("id")
	eventID, err := strconv.ParseInt(eventId, 10, 64)
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
	userID := context.GetInt64("userID")

if event.UserID != userID {
		context.JSON(http.StatusBadRequest, gin.H{ "message": "Not authorized to update"})
		return;
	}

	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "Event deleted successfully", "event": event})

}
