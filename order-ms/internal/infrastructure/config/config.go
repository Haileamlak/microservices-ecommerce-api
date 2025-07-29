package config

import (
    "log"
    "os"
	"github.com/joho/godotenv"
)

type Config struct {
    MongoURI      string
    MongoDatabase string
    GRPCPort      string
	HTTPPort      string
	ProductServiceURL string
	UserServiceURL string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

    cfg := &Config{
        MongoURI:      os.Getenv("MONGO_URI"),
        MongoDatabase: os.Getenv("MONGO_DB"),
        GRPCPort:      os.Getenv("GRPC_PORT"),
        HTTPPort:      os.Getenv("HTTP_PORT"),
		ProductServiceURL: os.Getenv("PRODUCT_SERVICE_URL"),
		UserServiceURL: os.Getenv("USER_SERVICE_URL"),
    }

    if cfg.MongoURI == "" || cfg.MongoDatabase == "" || cfg.GRPCPort == "" {
        log.Fatal("missing required environment variables")
    }

    return cfg
}
