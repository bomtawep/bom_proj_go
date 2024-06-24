package middleware

import (
	"bom_proj_go/pkg/common/configs"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

var authMiddleWare fiber.Handler

func Init() {
	env := configs.GetConfig()
	authMiddleWare = jwtware.New(jwtware.Config{
		SigningKey:     jwtware.SigningKey{Key: []byte(env.Secret)},
		SuccessHandler: authSuccess,
		ErrorHandler:   authError,
		TokenLookup:    "header:Authorization",
		AuthScheme:     "Bearer",
	})
}

func authError(c *fiber.Ctx, e error) error {
	err := c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
		"msg":   e.Error(),
	})
	if err != nil {
		return err
	}
	return nil
}

func authSuccess(c *fiber.Ctx) error {
	err := c.Next()
	if err != nil {
		return err
	}
	return nil
}

func GetAuthMiddleWare() fiber.Handler {
	return authMiddleWare
}
