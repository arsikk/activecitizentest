package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Postgres struct {
		DBHost     string `mapstructure:"DB_HOST"`
		DBPort     int    `mapstructure:"DB_PORT"`
		DBUser     string `mapstructure:"DB_USER"`
		DBPassword string `mapstructure:"DB_PASSWORD"`
		DBName     string `mapstructure:"DB_NAME"`
	}
	Redis struct {
		ADDR     string `json:"ADDR"`
		PASSWORD string `json:"PASSWORD"`
		DB       int    `json:"DB"`
	}
}

func LoadConfig() (config *Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		log.Fatal("config error")
	}

	return config, nil
}
