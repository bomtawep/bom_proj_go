package main

import (
	"bom_proj_go/api/routes"
	"bom_proj_go/configs"
	"bom_proj_go/internal/database"
	"bom_proj_go/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	configs.LoadEnv()
	database.ConnectDB()
	app := fiber.New()
	app.Use(cors.New())

	routes.GroupRoute(app)
	middleware.JWTProtected()

	err := app.Listen(`:` + configs.GetEnv("PORT"))
	if err != nil {
		log.Fatal("Error running server: ", err)
	}
	log.Println("Server running on port :" + configs.GetEnv("PORT"))
}
