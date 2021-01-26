package db

import (
	"context"
	"fl-auth/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DB_NAME         = "fl-auth"
	USER_COLLECTION = "users"
)

func (db *Client) CheckUser(lookupEmail string) (bool, error) {
	dbc := db.conn
	collection := dbc.Database(DB_NAME).Collection(USER_COLLECTION)

	// Checking if user exists
	var u models.User
	findErr := collection.FindOne(context.TODO(), bson.M{"email": lookupEmail}).Decode(&u)
	if findErr != nil && findErr != mongo.ErrNoDocuments {
		return false, findErr
	}
	if u.Email == lookupEmail {
		return true, nil
	}

	return false, nil
}

/* Function to add user for the client class. */
func (db *Client) AddUser(user models.User) error {
	dbc := db.conn
	collection := dbc.Database(DB_NAME).Collection(USER_COLLECTION)

	// Input validation
	if inputErr := user.IsValid(); inputErr != nil {
		return inputErr
	}

	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

/* Function to get all of the users for the client class. */
func (db *Client) GetUsersAll() ([]models.User, error) {
	dbc := db.conn
	var users []models.User

	collection := dbc.Database(DB_NAME).Collection(USER_COLLECTION)
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

/* Function to get a document by the email. */
func (db *Client) GetUserByEmail(lookupEmail string) (models.User, error) {
	dbc := db.conn
	collection := dbc.Database(DB_NAME).Collection(USER_COLLECTION)

	// Query db for user
	var u models.User
	findErr := collection.FindOne(context.TODO(), bson.M{"email": lookupEmail}).Decode(&u)
	if findErr != nil {
		return u, findErr
	}

	return u, nil
}

// func (db *Client) GetUserByName(name string) (Users, error) {
// 	dbc := db.conn
// 	user := Users{}

// 	return user, nil
// }
