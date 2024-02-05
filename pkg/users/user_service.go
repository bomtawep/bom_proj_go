package users

import (
	"bom_proj_go/pkg/common/configs"
	"bom_proj_go/pkg/common/database"
	"bom_proj_go/pkg/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var user models.User
var UserName models.UserName
var login models.Login

func insertUsers(context *fiber.Ctx) (*mongo.InsertOneResult, error) {
	// Declare all data
	userCollection := database.GetCollection("users")
	ctx, cancel := configs.CtxWithTimout()
	defer cancel()

	// Validate user input
	if err := context.BodyParser(&user); err != nil {
		return nil, err
	}
	if validationErr := validate.Struct(&user); validationErr != nil {
		return nil, validationErr
	}

	// Validate if username already exists
	resultUser := userCollection.FindOne(ctx, bson.M{"username": user.Username})
	err := resultUser.Decode(&user)
	if err == nil {
		return nil, fiber.NewError(400, "Username already exists")
	}
	log.Debug("resultUser", user)

	// Assign NewObjectID to user
	user.Id = primitive.NewObjectID()

	// Insert user to database
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
		Username:  user.Username,
		Password:  user.Password,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
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

func getUserByUsername(context *fiber.Ctx, user models.User) (*mongo.SingleResult, error) {
	userCollection := database.GetCollection("users")
	ctx, cancel := configs.CtxWithTimout()
	defer cancel()

	if err := context.BodyParser(&user); err != nil {
		return nil, err
	}
	if validationErr := validate.Struct(&user); validationErr != nil {
		return nil, validationErr
	}
	result := userCollection.FindOne(ctx, bson.M{"username": user.Username})
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return result, nil
}
