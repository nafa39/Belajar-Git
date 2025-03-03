package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RajaOngkirAPIKey string
	RajaOngkirAPIURL string
}

func Loadconfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		RajaOngkirAPIKey: os.Getenv("RAJAONGKIR_API_KEY"),
		RajaOngkirAPIURL: os.Getenv("RAJAONGKIR_API_URL"),
	}
}
