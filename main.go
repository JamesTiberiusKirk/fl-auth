package main

import (
	"fl-auth/config"
	"fl-auth/db"
	"fl-auth/server"

	"fmt"
)

func main() {

	env := config.GetEnv()

	dbClient, dbConnErr := db.Connect(env)
	if dbConnErr != nil {
		fmt.Println(dbConnErr)
	}

	server := server.Init(dbClient)
	server.Start(env.SERVER_PORT)
}
