package users

import (
	"bom_proj_go/configs"
	"bom_proj_go/internal/database"
	"bom_proj_go/models/users"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var userCollection *mongo.Collection = database.GetCollection(database.Database, "users")
var validate = validator.New()
var user users.User

func insertUsers(context *fiber.Ctx) (*mongo.InsertOneResult, error) {
	ctx, cancel := configs.CtxWithTimout()
	defer cancel()
	newUser := users.User{
		Id:        primitive.NewObjectID(),
		Username:  user.Username,
		Password:  user.Password,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
	}

	if err := context.BodyParser(&user); err != nil {
		return nil, err
	}
	if validationErr := validate.Struct(&user); validationErr != nil {
		return nil, validationErr
	}

	result, insertError := userCollection.InsertOne(ctx, newUser)
	if insertError != nil {
		panic("Error inserting user")
	}
	return result, insertError
}

func getUsers() ([]users.User, error) {
	ctx, cancel := configs.CtxWithTimout()
	defer cancel()
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var user []users.User
	if err = cursor.All(ctx, &user); err != nil {
		log.Fatal(err)
	}
	return user, err
}

func getUser(context *fiber.Ctx) (users.User, error) {
	hexUserId := context.Params("userId")
	ctx, cancel := configs.CtxWithTimout()
	defer cancel()
	objUserId, _ := primitive.ObjectIDFromHex(hexUserId)
	result := userCollection.FindOne(ctx, bson.M{"id": objUserId})
	log.Println(result)
	err := result.Decode(&user)
	if err != nil {
		return users.User{}, err
	}
	return user, nil
}

func updateUser(context *fiber.Ctx) (users.User, error) {
	hexUserId := context.Params("userId")
	ctx, cancel := configs.CtxWithTimout()
	defer cancel()
	if err := context.BodyParser(&user); err != nil {
		return user, err
	}
	if validationErr := validate.Struct(&user); validationErr != nil {
		return user, validationErr
	}
	objUserId, _ := primitive.ObjectIDFromHex(hexUserId)
	updateData := users.User{
		Id:        objUserId,
		Username:  user.Username,
		Password:  user.Password,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
	}
	result, err := userCollection.UpdateByID(ctx, bson.M{"id": objUserId}, bson.M{"$set": updateData})
	log.Println(result)
	if err != nil {
		return users.User{}, err
	}
	return user, err
}
