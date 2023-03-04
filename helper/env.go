package helper

import (
	"os"

	"github.com/joho/godotenv"
)

func InitializeEnv(pathToEnvFile string) {
	err := godotenv.Load(pathToEnvFile)
	if err != nil {
		panic(err)
	}

	if os.Getenv("DB_URL") == "" {
		panic("DB_URL env must be set")
	}

	if os.Getenv("DB_NAME") == "" {
		panic("DB_NAME env must be set")
	}
}
