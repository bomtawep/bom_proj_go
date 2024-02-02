package database

import (
	"bom_proj_go/pkg/common/configs"
	"context"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var conn *mongo.Client

func InitialDatabase() {
	env := configs.GetConfig()
	log.Info("Connecting to MongoDB...")
	clientOptions := options.Client().ApplyURI(env.MongoUri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Can't test ping database: ", err)
	}
	conn = client
	log.Info("Connected to MongoDB")
}
func GetCollection(collectionName string) *mongo.Collection {
	collection := conn.Database("bom-db").Collection(collectionName)
	return collection
}
func DisconnectDatabase() {
	err := conn.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Disconnected from MongoDB")
}
