package main

import (
	"bom_proj_go/pkg/common/configs"
	"bom_proj_go/pkg/common/database"
	"bom_proj_go/pkg/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func main() {
	env, err := configs.LoadEnv()
	fmt.Println("Config: ", env.MongoUrl)
	if err != nil {
		return
	}
	database.Initial()
	defer database.Disconnect()

	app := fiber.New()

	routes.GroupRoute(app)

	err = app.Listen(strconv.Itoa(env.Port))
	fmt.Println("Server is running on port", env.Port)
	if err != nil {
		log.Fatal("Error running server: ", err)
	}
}
