package config

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

/*
	GetEnvVariable - Return an environment variable, defined in a .env file, given a key.
*/
func GetEnvVariable(key string) string {
	if err := godotenv.Load(); err != nil {
		zap.S().Error("Cant read .env file")
	}

	return os.Getenv(key)
}
