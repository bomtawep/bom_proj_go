package upload

import (
	"bom_proj_go/pkg/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func CreateFile(context *fiber.Ctx) error {
	user, err := insertFile(context)
	if err != nil {
		return context.Status(
			http.StatusBadRequest).JSON(
			models.UserResponse{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Data:    &fiber.Map{"message": err.Error()},
			})
	}
	return context.Status(
		http.StatusCreated).JSON(
		models.UserResponse{
			Status:  http.StatusCreated,
			Message: http.StatusText(http.StatusCreated),
			Data:    &fiber.Map{"user": user},
		})
}

func GetFiles(context *fiber.Ctx) error {
	users, err := getFiles()
	if err != nil {
		return context.Status(
			http.StatusBadRequest).JSON(
			models.UserResponse{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Data:    &fiber.Map{"message": err.Error()},
			})
	}
	return context.Status(
		http.StatusCreated).JSON(
		models.UserResponse{
			Status:  http.StatusCreated,
			Message: http.StatusText(http.StatusCreated),
			Data:    &fiber.Map{"user": users},
		})
}
