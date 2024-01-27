package database

import (
	"bom_proj_go/pkg/common/configs"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var conn *mongo.Client
var cfg = configs.GetConfig()

func Initial() *mongo.Client {
	fmt.Println("Connecting to MongoDB")
	fmt.Println("Configs: ", cfg.Port)
	clientOptions := options.Client().ApplyURI(cfg.MongoUrl)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Config: ", configs.GetConfig())
	log.Println("Connected to MongoDB")
	return client
}

var Database = Initial()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("bom-db").Collection(collectionName)
	return collection
}
func Disconnect() {
	err := conn.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Disconnected from MongoDB")
}
