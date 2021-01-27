package db

import (
	"context"
	"fl-auth/config"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	Conn *mongo.Client
}

func Connect(env config.Env) (*Client, error) {
	mongoUri := "mongodb://" + env.DB_USER + ":" + env.DB_PASSWORD + "@" + env.DB_HOST + ":" + env.DB_PORT + "/" + env.DB_NAME
	fmt.Printf("Connecting to %s\n", mongoUri)

	// c, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	opts := options.Client().ApplyURI(mongoUri)
	c, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		return &Client{}, err

	}

	pingError := c.Ping(context.TODO(), nil)
	if pingError != nil {
		return &Client{}, pingError
	}

	client := &Client{
		Conn: c,
	}

	return client, nil
}
