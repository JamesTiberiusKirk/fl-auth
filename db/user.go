package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type Users struct {
	Email    string
	Username string
	Password string
}

/* Function to add user for the client class. */
func (db *Client) AddUser(user Users) error {
	dbc := db.conn

	collection := dbc.Database("fl-auth").Collection("users")

	fmt.Println(user)
	insertRes, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		return err
	}

	fmt.Println(insertRes)
	return nil
}

/* Function to get all of the users for the client class. */
func (db *Client) GetUsersAll() ([]Users, error) {
	dbc := db.conn
	var users []Users

	collection := dbc.Database("fl-auth").Collection("users")
	curr, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return users, err
	}

	err = curr.All(context.TODO(), &users)
	if err != nil {
		return users, err
	}

	return users, nil
}

// func (db *Client) GetUserById(id string) (Users, error) {
// 	dbc := db.conn
// 	user := Users{}

// 	return user, nil
// }

// func (db *Client) GetUserByEmail(email string) (Users, error) {
// 	dbc := db.conn
// 	user := Users{}

// 	return user, nil
// }

// func (db *Client) GetUserByName(name string) (Users, error) {
// 	dbc := db.conn
// 	user := Users{}

// 	return user, nil
// }
