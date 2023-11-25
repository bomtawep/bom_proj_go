package main

import (
	"bom_proj_go/api/routes"
	"bom_proj_go/internal/database"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	database.ConnectDB()
	routes.UserRoute(router)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal("Error running server: ", err)
	}
	log.Fatalln("Server running on port 8080: ", nil)
}
