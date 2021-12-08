package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// GetEnvVariable - Return an environment variable, defined in a .env file, given a key.
func GetEnvVariable(key string) (string, error) {
	if err := godotenv.Load(); err != nil {
		zap.S().Error("Cant read .env file")
		return "", err
	}
	envkey := os.Getenv(key)

	if envkey == "" {
		zap.S().Errorf("Key %s not found", envkey)
		return "", errors.New("key not found")
	}
	return envkey, nil
}
