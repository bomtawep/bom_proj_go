package models

type Config struct {
	Port     int    `env:"PORT" envDefault:"3000"`
	MongoUri string `env:"MONGO_URI" envDefault:"mongodb://localhost:27017"`
	Secret   string `env:"JWT_SECRET_KEY"`
}
