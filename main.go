package main

import (
	"fl-auth/config"
	"fl-auth/db"
	"fl-auth/server"
	"log"
)

func main() {

	env := config.GetEnv()

	dbClient, dbConnErr := db.Connect(env)
	if dbConnErr != nil {
		log.Fatalln(dbConnErr)
	}

	server := server.Init(dbClient)

	server.Start(env.SERVER_PORT)
}
