package auth

import (
	"bom_proj_go/pkg/common/configs"
	"bom_proj_go/pkg/models"
	"bom_proj_go/pkg/users"
	"bom_proj_go/pkg/utilities"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type (
	MsgLogin models.Login
	MsgToken models.Token
)

var validate = validator.New()
var login models.Login

func Login(context *fiber.Ctx) error {
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
	token, err := createToken(user.Id)
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
			Data:    &fiber.Map{"token": token},
		})
}

func GetSession(context *fiber.Ctx) error {
	userContext := context.Locals("user").(*jwt.Token)
	claims := userContext.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	return context.Status(
		http.StatusOK).JSON(
		models.UserResponse{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    &fiber.Map{"user_id": userId},
		})
}

func Logout(context *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), //Sets the expiry time an hour ago in the past.
		HTTPOnly: true,
	}

	context.Cookie(&cookie)
	return context.Status(
		http.StatusOK).JSON(
		models.UserResponse{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    &fiber.Map{"message": "Logged out"},
		})
}

func createToken(userId primitive.ObjectID) (string, error) {
	env := configs.GetConfig()

	claims := jwt.MapClaims{
		"id":    userId.Hex(),
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(env.Secret))
	if err != nil {
		return "", err
	}
	return t, nil
}
