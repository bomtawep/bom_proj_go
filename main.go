package main

import (
	"bom_proj_go/cmd/users"
	"bom_proj_go/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	database.InitializeDB()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running",
		})
	})

	http.HandleFunc("/create-user", users.CreateUser)

	err := r.Run()
	if err != nil {
		return
	}
}
