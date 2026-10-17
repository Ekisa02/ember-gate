package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDSN     string
	Port      string
	JWTSecret string
	LogLevel  string
}

func Load() (*Config, error) {
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.AutomaticEnv()

	cfg := &Config{
		DBDSN:     viper.GetString("DB_DSN"),
		Port:      viper.GetString("PORT"),
		JWTSecret: viper.GetString("JWT_SECRET"),
		LogLevel:  viper.GetString("LOG_LEVEL"),
	}

	return cfg, nil
}
