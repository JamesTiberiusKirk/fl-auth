package main

import (
	"fmt"

	"fl-auth/config"
	"fl-auth/server"
)

func main() {
	env, err := config.GetEnv()

	if err != nil {
		fmt.Println(err.Error())
	}

	server := server.Init()

	server.Start(env.SERVER_PORT)
}
