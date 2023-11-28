package utils

import (
	dotenv "github.com/joho/godotenv"
	"path"
)

func LoadEnv() {
	if err := dotenv.Load(path.Join(GetCurrentDir(), "..", ".env")); err != nil {
		panic("Error loading .env file")
	}
}
