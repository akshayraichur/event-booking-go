package routes

import (
	"net/http"

	"akshayraichur.com/event-booking-go/models"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Could not save user"})
	}

	context.JSON(http.StatusCreated, gin.H{"status": "User created successfully"})
	
}