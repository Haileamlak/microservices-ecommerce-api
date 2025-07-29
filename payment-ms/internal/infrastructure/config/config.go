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
	StripeSecretKey string
	UserServiceURL string
	OrderServiceURL string
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
		StripeSecretKey: os.Getenv("STRIPE_SECRET_KEY"),
		UserServiceURL: os.Getenv("USER_SERVICE_URL"),
		OrderServiceURL: os.Getenv("ORDER_SERVICE_URL"),
    }

    if cfg.MongoURI == "" || cfg.MongoDatabase == "" || cfg.GRPCPort == "" || cfg.StripeSecretKey == "" {
        log.Fatal("missing required environment variables")
    }

    return cfg
}
