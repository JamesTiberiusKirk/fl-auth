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

// CheckUser function for checking if user exists based on email.
func (db *Client) CheckUser(lookupEmail string) (bool, error) {
	dbc := db.Conn
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

// AddUser function to add user for the client class.
func (db *Client) AddUser(user models.User) error {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(USER_COLLECTION)

	// Input validation
	if inputErr := user.IsValid(); inputErr != nil {
		return inputErr
	}

	// Inserting manually because I need mongo to autogenerate the ID
	insert := bson.M{
		"email":    user.Email,
		"username": user.Username,
		"password": user.Password,
		"roles":    user.Roles,
	}

	_, err := collection.InsertOne(context.TODO(), insert)
	if err != nil {
		return err
	}

	return nil
}

// GetUsersAll function to get all of the users for the client class.
func (db *Client) GetUsersAll() ([]models.User, error) {
	dbc := db.Conn
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

// GetUserByEmail function to get a document by the email.
func (db *Client) GetUserByEmail(lookupEmail string) (models.User, error) {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(USER_COLLECTION)

	// Query db for user
	var u models.User
	filter := bson.M{"email": lookupEmail}
	findErr := collection.FindOne(context.TODO(), filter).Decode(&u)

	if findErr != nil {
		return u, findErr
	}

	return u, nil
}
