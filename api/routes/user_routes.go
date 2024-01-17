package routes

import (
	"bom_proj_go/pkg/users"
	"github.com/gofiber/fiber/v2"
)

func userGroup(group fiber.Router) {
	userGroup := group.Group("/users")
	userGroup.Post("/", users.CreateUser)
	userGroup.Get("/:userId", users.GetUser)
	userGroup.Put("/:userId", users.UpdateUser)
	userGroup.Delete("/:userId", users.DeleteUser)
	userGroup.Get("/", users.GetUsers)
}
