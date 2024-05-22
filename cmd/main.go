package main

import (
	"bom_proj_go/pkg/common/configs"
	"bom_proj_go/pkg/common/database"
	"bom_proj_go/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"strconv"
)

func main() {
	// Load environment variables
	env := configs.LoadEnv()

	// Connect to MongoDB
	database.InitialDatabase()
	defer database.DisconnectDatabase()

	// Create a new Fiber app
	app := fiber.New()

	app.Static("/uploads", "./uploads")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	routes.GroupRoute(app)
	err := app.Listen(`:` + strconv.Itoa(env.Port))
	if err != nil {
		log.Fatal("Fiber app server error: ", err)
	}
}
