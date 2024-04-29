package main

import (
	"fmt"

	"github.com/chethanbhat/go-rest-api/db"
	"github.com/chethanbhat/go-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	fmt.Println("Welcome to Rest API")
	server := gin.Default()

	// Register Routes
	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost

}
