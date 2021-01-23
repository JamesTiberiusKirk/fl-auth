// main.go
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	SERVER_PORT string
}

func GetEnv() Env {

	envConfig := Env{}

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("No .env file")
	} else {
		fmt.Print("Loading from .env..")
	}

	envConfig.DB_HOST = os.Getenv("DB_HOST")
	envConfig.DB_NAME = os.Getenv("DB_NAME")
	envConfig.DB_PORT = os.Getenv("DB_PORT")
	envConfig.DB_USER = os.Getenv("DB_USER")
	envConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	envConfig.SERVER_PORT = os.Getenv("SERVER_PORT")

	return envConfig
}
