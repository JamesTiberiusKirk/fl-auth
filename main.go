package main

import (
	"fl-auth/config"
	"fl-auth/db"
	"fl-auth/server"
)

func main() {

	env := config.GetEnv()

	dbClient := db.Connect(env)
	server := server.Init(dbClient)

	server.Start(env.SERVER_PORT)
}
