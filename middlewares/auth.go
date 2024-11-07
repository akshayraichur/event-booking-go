package middlewares

import (
	"fmt"
	"net/http"

	"akshayraichur.com/event-booking-go/utils"
	"github.com/gin-gonic/gin"
)


func Authenticate(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")

		fmt.Println("Token: ", token)

		if token == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		userID, err := utils.VerifyToken(token)
		if err != nil {
			fmt.Println("Error: ", err)
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// callbackFunction(context, userID)
		context.Set("userID", userID)
		context.Next()
	}