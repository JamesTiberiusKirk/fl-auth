package db

import (
	"context"
	"fl-auth/config"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client contains the mongo connection and custom functions.
type Client struct {
	Conn *mongo.Client
}

// Connect for connecting to the mongo serer.
func Connect(env config.Env) (*Client, error) {
	mongoURI := "mongodb://" + env.DB_USER + ":" + env.DB_PASSWORD + "@" + env.DB_HOST + ":" + env.DB_PORT + "/" + env.DB_NAME
	fmt.Printf("Connecting to %s\n", mongoURI)

	opts := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	c, err := mongo.Connect(ctx, opts)
	defer cancel()

	if err != nil {
		fmt.Println("Couldn't connect to the database", err)
		return &Client{}, err
	} else {
		fmt.Println("Database connected!")
	}

	pingError := c.Ping(context.TODO(), nil)
	if pingError != nil {
		return &Client{}, pingError
	} else {
		fmt.Println("Database ping")
	}

	client := &Client{
		Conn: c,
	}

	return client, nil
}
