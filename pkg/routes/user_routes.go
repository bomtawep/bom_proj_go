package routes

import (
	"bom_proj_go/pkg/auth"
	"bom_proj_go/pkg/upload"
	"bom_proj_go/pkg/users"
	"github.com/gofiber/fiber/v2"
)

func userGroup(group fiber.Router) {
	userGroup := group.Group("/users")
	userGroup.Post("/", users.CreateUser)
	userGroup.Post("/image", upload.CreateFile)
	userGroup.Post("/email", users.GetUserByEmail)
	group.Use(auth.AuthorizationRequired)
	userGroup.Get("/image", upload.GetFiles)
	userGroup.Get("/:userId", users.GetUser)
	userGroup.Put("/:userId", users.UpdateUser)
	userGroup.Delete("/:userId", users.DeleteUser)
	userGroup.Get("/", users.GetUsers)
}
