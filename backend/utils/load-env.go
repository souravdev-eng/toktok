package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load(".env")
	return err
}

func GetEnvValue(key string) string {
	return os.Getenv(key)
}
