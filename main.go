package main

import (
	"fl-auth/config"
	"fl-auth/server"
)

func main() {
	env := config.GetEnv()

	server := server.Init()

	server.Start(env.SERVER_PORT)
}
