package main

import (
	"bom_proj_go/pkg/common/configs"
	"bom_proj_go/pkg/common/database"
	"bom_proj_go/pkg/common/middleware/middleware"
	"bom_proj_go/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"strconv"
)

func main() {
	env := configs.LoadEnv()
	middleware.Init()
	database.InitialDatabase()
	defer database.DisconnectDatabase()

	app := fiber.New()

	app.Static("/uploads", "./uploads")
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "*",
		AllowCredentials: true,
	}))
	routes.GroupRoute(app)
	err := app.Listen(`:` + strconv.Itoa(env.Port))
	if err != nil {
		log.Fatal("Fiber app server error: ", err)
	}
}
