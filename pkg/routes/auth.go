package routes

import (
	"bom_proj_go/pkg/auth"
	"github.com/gofiber/fiber/v2"
)

func authGroup(group fiber.Router) {
	authGroup := group.Group("/auth")
	authGroup.Post("/login", auth.Auth)
	authGroup.Get("/session", auth.GetSession)
	authGroup.Post("/logout", auth.Logout)
}
