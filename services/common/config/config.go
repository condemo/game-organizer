package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var EnvConfig = newNevConfig()

type envConfig struct {
	Host string
	Port string
	Name string
	User string
	Pass string
}

func newNevConfig() *envConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	return &envConfig{
		Host: os.Getenv("POSTGRES_HOST"),
		Port: os.Getenv("POSTGRES_PORT"),
		Name: os.Getenv("POSTGRES_DB_NAME"),
		User: os.Getenv("POSTGRES_USER"),
		Pass: os.Getenv("POSTGRES_PASS"),
	}
}
