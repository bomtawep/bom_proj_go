package configs

import (
	"bom_proj_go/pkg/models"
	"fmt"
	"github.com/02amanag/environment"
)

var configs = models.Config{}

func LoadEnv() (models.Config, error) {
	cfg := &configs
	err := environment.Unmarshal(*cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", *cfg) //output -> testUser
	return configs, err
}

func GetConfig() models.Config {
	return configs
}
