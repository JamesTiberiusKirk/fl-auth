package db

import (
	"context"
	"fl-auth/config"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	conn *mongo.Client
}

func Connect(env config.Env) *Client {
	mongoUri := "mongodb://" + env.DB_USER + ":" + env.DB_PASSWORD + "@" + env.DB_HOST + ":" + env.DB_PORT + "/" + env.DB_NAME
	fmt.Printf("Connecting to %s\n", mongoUri)

	// c, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	opts := options.Client().ApplyURI(mongoUri)
	c, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Fatal(err)
	}

	client := &Client{
		conn: c,
	}

	return client
}
