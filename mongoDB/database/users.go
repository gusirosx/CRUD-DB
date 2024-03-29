package database

import (
	"context"
	"crudAPI/types"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create an unexported global variable to hold the collection connection pool.
func (client *MongoClient) getCollection() (collection *mongo.Collection) {
	return client.OpenCollection("user")
}

// Get all users from the DB by its id
func (client *MongoClient) GetUsers(ctx *gin.Context) (primitive.M, error) { //([]types.User, error)

	// Get all users from the DB by its id
	recordPerPage, err := strconv.Atoi(ctx.Query("recordPerPage"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}

	startIndex, _ := strconv.Atoi(ctx.Query("startIndex"))
	matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}}
	groupStage := bson.D{{Key: "$group", Value: bson.D{
		{Key: "_id", Value: bson.D{{Key: "_id", Value: "null"}}},
		{Key: "total_count", Value: bson.D{{Key: "$sum", Value: 1}}},
		{Key: "data", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}}}}}
	projectStage := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "total_count", Value: 1},
			{Key: "user_items", Value: bson.D{
				{Key: "$slice", Value: []interface{}{"$data", startIndex, recordPerPage}}}}}}}

	var queryCtx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// get mongo collection
	collection := client.getCollection()
	//var response *mongo.Cursor
	response, err := collection.Aggregate(queryCtx, mongo.Pipeline{matchStage, groupStage, projectStage})
	if err != nil {
		log.Println(err.Error())
		// return empty users slice on error
		return primitive.M{}, nil
	}
	// create an empty list of type []bson.M
	var usersList []bson.M
	if err = response.All(ctx, &usersList); err != nil {
		log.Println(err.Error())
		// return empty users slice on error
		return primitive.M{}, nil
	}
	// send the response message
	return usersList[0], nil
}

// Get one user from the DB by its id
func (client *MongoClient) GetUser(UID string) (types.User, error) {
	var user types.User
	var queryCtx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// Get a primitive ObjectID from a hexadecimal string
	idPrimitive, err := primitive.ObjectIDFromHex(UID)
	if err != nil {
		log.Println(err.Error())
		return types.User{}, err
	}

	// get mongo collection
	collection := client.getCollection()
	// Call the FindOne() method by passing BSON
	if err := collection.FindOne(queryCtx, bson.M{"_id": idPrimitive}).Decode(&user); err != nil {
		return types.User{}, err
	}
	return user, nil
}

// Create one user into DB
func (client *MongoClient) CreateUser(user types.User) (types.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// get a unique userID
	user.ID = primitive.NewObjectID()
	// get mongo collection
	collection := client.getCollection()
	//Insert one entry
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Println(err.Error())
		return types.User{}, fmt.Errorf("unable to create user")
	}
	return user, nil
}

// Update one user from the DB by its id
func (client *MongoClient) UpdateUser(id string, user types.User) error {

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
	// get mongo collection
	collection := client.getCollection()
	_, err = collection.UpdateByID(ctx, idPrimitive, update, opt)
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("unable to update user")
	}

	return nil
}

// Delete one user from the DB by its id
func (client *MongoClient) DeleteUser(id string) error {
	var queryCtx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// Declare a primitive ObjectID from a hexadecimal string
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	// get mongo collection
	collection := client.getCollection()
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
