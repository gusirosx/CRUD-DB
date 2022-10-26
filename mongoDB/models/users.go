package models

import (
	"context"
	"crudAPI/database"
	"crudAPI/types"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create an unexported global variable to hold the database connection pool.
var client *database.MongoClient = database.MongoInstance()

// Create an unexported global variable to hold the collection connection pool.
var collection *mongo.Collection = client.OpenCollection("user")

// Create one user into DB
func CreateUser(user types.User) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// get a unique userID
	user.Id = primitive.NewObjectID()
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("unable to create user")
	}

	return nil
}

// Update one user from the DB by its id
func UpdateUser(id string, user types.User) error {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	// Declare a primitive ObjectID from a hexadecimal string
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	var updatedUser primitive.D
	updatedUser = append(updatedUser,
		bson.E{Key: "name", Value: user.Name},
		bson.E{Key: "gender", Value: user.Gender},
		bson.E{Key: "age", Value: user.Age})
	opt := options.Update().SetUpsert(true)
	update := bson.D{{Key: "$set", Value: updatedUser}}
	_, err = collection.UpdateByID(ctx, idPrimitive, update, opt)
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("unable to update user")
	}

	return nil
}

// Delete one user from the DB by its id
func DeleteUser(id string) error {
	var queryCtx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// Declare a primitive ObjectID from a hexadecimal string
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// Call the DeleteOne() method by passing BSON
	res, err := collection.DeleteOne(queryCtx, bson.M{"_id": idPrimitive})
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("unable to delete user")
	} else if res.DeletedCount == 0 {
		return fmt.Errorf("there is no such user for be deleted")
	}
	return nil
}
