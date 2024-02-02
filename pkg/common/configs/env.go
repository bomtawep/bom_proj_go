package configs

import (
	"bom_proj_go/pkg/models"
	"github.com/02amanag/environment"
	"github.com/gofiber/fiber/v2/log"
)

var configs = models.Config{}

func LoadEnv() models.Config {
	err := environment.Unmarshal(&configs)
	if err != nil {
		log.Fatal("Error loading environment variables: ", err)
	}
	return configs
}

func GetConfig() models.Config {
	return configs
}
