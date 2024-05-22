package routes

import (
	"github.com/gofiber/fiber/v2"
)

func GroupRoute(app *fiber.App) {
	group := app.Group("/api")
	authGroup(group)
	userGroup(group)
}
