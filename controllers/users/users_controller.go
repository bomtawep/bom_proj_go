package users

import (
	"bom_proj_go/internal/database"
	"bom_proj_go/models/users"
	"bom_proj_go/responses"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.GetCollection(database.Database, "users")
var validate = validator.New()

func CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user users.User
	defer cancel()

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newUser := users.User{
		Id:        primitive.NewObjectID(),
		Username:  user.Username,
		Password:  user.Password,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
	}

	// Validate if username already exists and insert if not
	var result users.User
	err := userCollection.FindOne(ctx, users.User{Username: user.Username}).Decode(&result)
	if err != nil {
		log.Println("User ", user.Username, " already exists")
		return c.Status(http.StatusConflict).JSON(responses.UserResponse{Status: http.StatusConflict, Message: "error", Data: &fiber.Map{"data": "Username already exists"}})
	}
	// Insert user
	insertResult, insertErr := userCollection.InsertOne(ctx, newUser)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	log.Println("Inserted user with ID: ", insertResult.InsertedID)

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": newUser}})
}
