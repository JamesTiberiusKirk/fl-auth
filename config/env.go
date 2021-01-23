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

func GetEnv() (Env, error) {

	envConfig := Env{}

	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("No .env file")
		// log.Fatalf("Error loading .env file")
		// return envConfig, err
	}

	envConfig.DB_HOST = os.Getenv("DB_HOST")
	envConfig.DB_NAME = os.Getenv("DB_NAME")
	envConfig.DB_PORT = os.Getenv("DB_PORT")
	envConfig.DB_USER = os.Getenv("DB_USER")
	envConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	envConfig.SERVER_PORT = os.Getenv("SERVER_PORT")

	fmt.Println(envConfig)

	return envConfig, nil

}
