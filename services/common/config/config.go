package config

import (
	"os"
)

var (
	EnvConfig = newNevConfig()
	IGDBToken string
)

type envConfig struct {
	Host string
	Port string
	Name string
	User string
	Pass string
}

func newNevConfig() *envConfig {
	return &envConfig{
		Host: os.Getenv("POSTGRES_HOST"),
		Port: os.Getenv("POSTGRES_PORT"),
		Name: os.Getenv("POSTGRES_DB_NAME"),
		User: os.Getenv("POSTGRES_USER"),
		Pass: os.Getenv("POSTGRES_PASS"),
	}
}
