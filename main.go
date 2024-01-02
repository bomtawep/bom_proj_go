package main

import (
	"bom_proj_go/api/routes"
	"bom_proj_go/configs"
	"bom_proj_go/internal/database"
	"bom_proj_go/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	configs.LoadEnv()
	app := fiber.New()

	database.ConnectDB()
	routes.GroupRoute(app)
	middleware.JWTProtected()
	app.Get("/", func(context *fiber.Ctx) error {
		return context.SendString("Hello, World!")
	})

	err := app.Listen(`:` + configs.GetEnv("PORT"))
	if err != nil {
		log.Fatal("Error running server: ", err)
	}
	log.Fatalln("Server running on port 8080: ", nil)
}
