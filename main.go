package main

import (
	"fl-auth/config"
	"fl-auth/db"
	"fmt"
)

func main() {

	env := config.GetEnv()

	dbClient := db.Connect(env)
	// server := server.Init(dbClient)

	// err := dbClient.AddUser()
	// if err != nil {
	// 	fmt.Println("err")
	// 	fmt.Println(err)
	// }

	data, e := dbClient.GetUsersAll()
	if e != nil {
		fmt.Println("err")
		fmt.Println(e)
	}

	for _, value := range data {
		fmt.Printf("- %s\n", value.Email)
	}

	fmt.Println(data)
	// server.Start(env.SERVER_PORT)
}
