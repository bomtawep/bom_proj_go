package users

import (
	"bom_proj_go/pkg/common/configs"
	"bom_proj_go/pkg/common/database"
	"bom_proj_go/pkg/models"
	"bom_proj_go/pkg/utilities"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var user models.User
var email models.Email

func insertUsers(context *fiber.Ctx) (*mongo.InsertOneResult, error) {
	userCollection := database.GetCollection("users")
	ctx, cancel := configs.CtxWithTimout()
	defer cancel()
	user.Id = primitive.NewObjectID()

	if err := context.BodyParser(&user); err != nil {
		return nil, err
	}
	if validationErr := validate.Struct(&user); validationErr != nil {
		return nil, validationErr
	}
	hash, err := utilities.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hash
	result, insertError := userCollection.InsertOne(ctx, &user)
	if insertError != nil {
		panic("Error inserting user")
	}
	return result, insertError
}

func getUsers() ([]models.User, error) {
	userCollection := database.GetCollection("users")
	ctx, cancel := configs.CtxWithTimout()
	defer cancel()
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var user []models.User
	if err = cursor.All(ctx, &user); err != nil {
		log.Fatal(err)
	}
	return user, err
}

func getUser(context *fiber.Ctx) (models.User, error) {
	userCollection := database.GetCollection("users")
	hexUserId := context.Params("userId")
	ctx, cancel := configs.CtxWithTimout()
	defer cancel()
	objUserId, _ := primitive.ObjectIDFromHex(hexUserId)
	result := userCollection.FindOne(ctx, bson.M{"id": objUserId})
	err := result.Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func updateUser(context *fiber.Ctx) (*mongo.UpdateResult, error) {
	userCollection := database.GetCollection("users")
	hexUserId := context.Params("userId")
	ctx, cancel := configs.CtxWithTimout()
	defer cancel()
	if err := context.BodyParser(&user); err != nil {
		return nil, err
	}
	if validationErr := validate.Struct(&user); validationErr != nil {
		return nil, validationErr
	}
	objUserId, _ := primitive.ObjectIDFromHex(hexUserId)
	updateData := models.User{
		Id:        objUserId,
		Email:     user.Email,
		Password:  user.Password,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Gender:    user.Gender,
		Phone:     user.Phone,
		Birthdate: user.Birthdate,
	}
	result, err := userCollection.UpdateOne(ctx, bson.M{"id": objUserId}, bson.M{"$set": updateData})
	if err != nil {
		return nil, err
	}
	return result, err
}

func deleteUser(context *fiber.Ctx) (string, error) {
	userCollection := database.GetCollection("users")
	hexUserId := context.Params("userId")
	ctx, cancel := configs.CtxWithTimout()
	defer cancel()
	objUserId, _ := primitive.ObjectIDFromHex(hexUserId)
	_, err := userCollection.DeleteOne(ctx, bson.M{"id": objUserId})
	if err != nil {
		return hexUserId, err
	}
	return hexUserId, err
}

func getUserByEmail(context *fiber.Ctx) (models.User, error) {
	userCollection := database.GetCollection("users")
	ctx, cancel := configs.CtxWithTimout()
	defer cancel()

	if err := context.BodyParser(&email); err != nil {
		return models.User{}, err
	}
	if validationErr := validate.Struct(&email); validationErr != nil {
		return models.User{}, validationErr
	}
	result := userCollection.FindOne(ctx, bson.M{"email": email.Email})
	err := result.Decode(&user)
	if !errors.Is(err, mongo.ErrNoDocuments) {
		return models.User{}, errors.New("User not found ")
	}
	return models.User{}, nil
}

func GetUserPassword(userName string) models.User {
	userCollection := database.GetCollection("users")
	ctx, cancel := configs.CtxWithTimout()
	defer cancel()

	result := userCollection.FindOne(ctx, bson.M{"email": userName})
	err := result.Decode(&user)
	if err != nil {
		return user
	}

	return user
}
