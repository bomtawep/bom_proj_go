package auth

import (
	"bom_proj_go/pkg/models"
	"bom_proj_go/pkg/users"
	"bom_proj_go/pkg/utilities"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

const (
	jwtSecret = "bomtawep"
)

type (
	MsgLogin models.Login
	MsgToken models.Token
)

var store = session.New()
var validate = validator.New()
var login models.Login

func Auth(context *fiber.Ctx) error {
	sess, err := store.Get(context)
	if err := context.BodyParser(&login); err != nil {
		return context.Status(
			http.StatusBadRequest).JSON(
			models.UserResponse{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Data:    &fiber.Map{"message": err.Error()},
			})
	}
	if validationErr := validate.Struct(&login); validationErr != nil {
		return context.Status(
			http.StatusBadRequest).JSON(
			models.UserResponse{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Data:    &fiber.Map{"message": validationErr.Error()},
			})
	}
	user := users.GetUserPassword(login.Username)
	if login.Username != user.Email || !utilities.CheckPasswordHash(login.Password, user.Password) {
		return context.Status(
			http.StatusUnauthorized).JSON(
			models.UserResponse{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
				Data:    &fiber.Map{"message": "Invalid username or password"},
			})
	}
	token, err := createToken()
	if err != nil {
		return context.Status(
			http.StatusInternalServerError).JSON(
			models.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Data:    &fiber.Map{"message": err.Error()},
			})
	}
	sess.Set("access_token", token.AccessToken)
	if err := sess.Save(); err != nil {
		return context.Status(
			http.StatusInternalServerError).JSON(
			models.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Data:    &fiber.Map{"message": err.Error()},
			})
	}
	return context.Status(
		http.StatusOK).JSON(
		models.UserResponse{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    &fiber.Map{"token": token},
		})
}

func GetSession(context *fiber.Ctx) error {
	sess, err := store.Get(context)
	if err != nil {
		return context.Status(
			http.StatusInternalServerError).JSON(
			models.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Data:    &fiber.Map{"message": err.Error()},
			})
	}
	return context.Status(
		http.StatusOK).JSON(
		models.UserResponse{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    &fiber.Map{"session": sess.Get("access_token")},
		})

}

func Logout(context *fiber.Ctx) error {
	sess, err := store.Get(context)
	if err != nil {
		return context.Status(
			http.StatusInternalServerError).JSON(
			models.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Data:    &fiber.Map{"message": err.Error()},
			})
	}
	err = sess.Destroy()
	if err != nil {
		return err
	}
	return context.Status(
		http.StatusOK).JSON(
		models.UserResponse{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    &fiber.Map{"message": "Logged out"},
		})

}

func createToken() (MsgToken, error) {
	var msgToken MsgToken
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = utils.UUID()
	claims["name"] = "bom"
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return msgToken, err
	}
	msgToken.AccessToken = t

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = utils.UUID()
	rtClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	rt, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return msgToken, err
	}
	msgToken.RefreshToken = rt
	return msgToken, nil
}

func AuthorizationRequired(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecret),
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			return authError(c, e)
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			return authSuccess(c)
		},
	})(c)
}

func authError(c *fiber.Ctx, e error) error {
	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
		"msg":   e.Error(),
	})
	return nil
}

func authSuccess(c *fiber.Ctx) error {
	err := c.Next()
	if err != nil {
		return err
	}
	return nil
}
