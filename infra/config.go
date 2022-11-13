package infra

import (
	"errors"
	"os"
)

type Config struct {
	MongoDBUri string
	Listen     string
}

func LoadConfig() (*Config, error) {
	result := &Config{}
	result.MongoDBUri = os.Getenv("MONGODB_URI")

	if result.MongoDBUri == "" {
		return nil, errors.New("MONGODB_URI is not set")
	}

	result.Listen = os.Getenv("LISTEN")
	if result.Listen == "" {
		result.Listen = ":8080"
	}

	return result, nil
}
