package main

import (
	"bom_proj_go/api/routes"
	"github.com/gofiber/fiber/v2"

	//"bom_proj_go/api/routes"
	"bom_proj_go/internal/database"
	"log"
)

func main() {
	app := fiber.New()

	database.ConnectDB()
	routes.GroupRoute(app)

	err := app.Listen(":8000")
	if err != nil {
		log.Fatal("Error running server: ", err)
	}
	log.Fatalln("Server running on port 8080: ", nil)
}
