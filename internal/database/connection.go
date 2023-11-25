package database

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func InitializeDB() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found")
	}
	uri := os.Getenv("MONGO_HOST")
	if uri == "" {
		log.Fatal("MONGO_HOST is not set")
	}
	var err error

	db, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer func() {
		if err := db.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}
