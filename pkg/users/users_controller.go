package users

import (
	"bom_proj_go/pkg/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func CreateUser(context *fiber.Ctx) error {
	user, err := insertUsers(context)
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

func GetUsers(context *fiber.Ctx) error {
	users, err := getUsers()
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

func GetUser(context *fiber.Ctx) error {
	users, err := getUser(context)
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

func UpdateUser(context *fiber.Ctx) error {
	users, err := updateUser(context)
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

func DeleteUser(context *fiber.Ctx) error {
	deleteResult, err := deleteUser(context)
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
			Data:    &fiber.Map{"deleted": deleteResult},
		})
}

func GetUserByEmail(context *fiber.Ctx) error {
	users, err := getUserByEmail(context)
	if err != nil {
		return context.Status(
			http.StatusBadRequest).JSON(
			models.UserResponse{
				Status:  http.StatusNotFound,
				Message: http.StatusText(http.StatusNotFound),
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
