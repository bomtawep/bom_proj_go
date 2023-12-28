package routes

import (
	"bom_proj_go/controllers/users"
	"github.com/gofiber/fiber/v2"
)

func userGroup(group fiber.Router) {

	userGroup := group.Group("/users")
	userGroup.Post("/", users.CreateUser)
	//app.GET("/user/:userId", users.GetAUser)
	//app.PUT("/user/:userId", users.EditAUser)
	//app.DELETE("/user/:userId", users.DeleteAUser)
	//app.GET("/users", users.GetAllUsers)
}
