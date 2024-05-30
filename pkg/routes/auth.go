package routes

import (
	"bom_proj_go/pkg/auth"
	"bom_proj_go/pkg/common/middleware/middleware"
	"github.com/gofiber/fiber/v2"
)

func authGroup(group fiber.Router) {
	authGroup := group.Group("/auth")
	authGroup.Post("/login", auth.Login)
	authGroup.Post("/logout", auth.Logout)
	authGroup.Use(middleware.GetAuthMiddleWare())
	authGroup.Get("/session", auth.GetSession)
}
