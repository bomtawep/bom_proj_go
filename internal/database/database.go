package database

import (
	"bom_proj_go/configs"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectDB() {
	clientOptions := options.Client().ApplyURI(configs.GetEnv("MONGO_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	conn = client
}

var conn *mongo.Client

func GetCollection(collectionName string) *mongo.Collection {
	collection := conn.Database("bom-db").Collection(collectionName)
	return collection
}
