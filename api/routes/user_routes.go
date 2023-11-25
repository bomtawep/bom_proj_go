package routes

import (
	"bom_proj_go/controllers/users"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", users.CreateUser())
	router.GET("/user/:userId", users.GetAUser())
	router.PUT("/user/:userId", users.EditAUser())
	router.DELETE("/user/:userId", users.DeleteAUser())
	router.GET("/users", users.GetAllUsers())
}
