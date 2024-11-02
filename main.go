package main

import (
	"akshayraichur.com/event-booking-go/db"
	"akshayraichur.com/event-booking-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default() // sets up an HTTP server with some default middleware (with logger and recovery middleware)

	routes.RegisterRoutes(server)
	
	server.Run(":8080") // listen and serve on

}
