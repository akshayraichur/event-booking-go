package routes

import (
	"net/http"

	"akshayraichur.com/event-booking-go/models"
	"akshayraichur.com/event-booking-go/utils"
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

func login(context *gin.Context){
	// login
	
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.Authenticate()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "message": "Unauthorized"})
		return
	}

	token , err := utils.GenerateToken(user.Email, string(user.ID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Could not generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User authenticated", "token": token})
}