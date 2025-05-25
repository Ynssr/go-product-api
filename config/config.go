package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
}

var AppConfig *Config

func LoadConfig() {
	// Proje kök dizinindeki .env dosyasını yükle
	envPath := filepath.Join("..", ".env") // cmd'den bir üst dizin
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("env file not load: %v", err)
	}

	AppConfig = &Config{MongoURI: getEnv("MONGO_URI")}
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s env variable not defined", key)
	}

	return value
}
