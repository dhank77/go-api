package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
}

func LoadConfig() (*Config, error) {
	// Read from environment variables
	viper.AutomaticEnv()
	viper.BindEnv("DATABASE_URL")

	// Read from .env file as fallback
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()

	var config Config
	config.DatabaseURL = viper.GetString("DATABASE_URL")

	return &config, nil
}
