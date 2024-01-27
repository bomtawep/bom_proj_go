package models

type Config struct {
	Port     int    `env:"PORT" envDefault:"3000"`
	MongoUrl string `env:"MONGO_URI" envDefault:"mongodb://localhost:27017"`
}
