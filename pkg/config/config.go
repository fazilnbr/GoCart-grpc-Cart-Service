package config

import (
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBCart     string `mapstructure:"DB_CART"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_CART", "DB_PORT", "DB_PASSWORD",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./pkg/config/envs/")
	viper.SetConfigFile("./pkg/config/envs/dev.env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}
