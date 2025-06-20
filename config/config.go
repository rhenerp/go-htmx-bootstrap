package config

import (
	"fmt"
	"os"
)

type Config struct {
	ENV string
	Port string
}

func LoadConfig() Config {
	env := os.Getenv("ENV")

fmt.Println(env)
	switch env {
	case "production":
		return LoadProdConfig()
	case "staging":
		return LoadStagingConfig()
	default:
		return LoadDevConfig()
	}
}

func GetEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	} else {
		return fallback
	}
}