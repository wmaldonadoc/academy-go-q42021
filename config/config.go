package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Cant read .env file")
	}
	return os.Getenv(key)
}
