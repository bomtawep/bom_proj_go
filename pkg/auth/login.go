package auth

import (
	"bom_proj_go/pkg/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

var validate = validator.New()
var login models.Login

func Login(c *fiber.Ctx) error {
	logins := models.Login{
		Username: login.Username,
		Password: login.Password,
	}
	log.Default().Println("body", logins.Username)

	if c.Params("password") != "john" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
func Logout(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
