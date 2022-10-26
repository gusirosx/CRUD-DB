package database

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

// // Get one user from the DB by its id
// func (db *SqlClient) GetUser(UID string) (types.User, error) {
// 	// create an empty user of type entity.User
// 	var user types.User
// 	// create the select sql query
// 	comand := "select * from users where id=$1"
// 	// execute the sql statement
// 	row := db.QueryRow(comand, UID)
// 	// unmarshal the row object to user struct
// 	if err := row.Scan(&user.ID, &user.Name, &user.Gender, &user.Age); err != nil {
// 		return types.User{}, err
// 	}
// 	return user, nil
// }

// // Create one user into DB
// func (db *SqlClient) CreateUser(user types.User) (types.User, error) {
// 	// get a unique userID
// 	user.ID = uuid.New().String()
// 	// execute the sql statement
// 	comand, err := db.Prepare("insert into users(id,name,gender,age) values($1, $2, $3, $4)")
// 	if err != nil {
// 		return types.User{}, fmt.Errorf("unable to create the user:" + err.Error())
// 	}
// 	defer comand.Close()
// 	comand.Exec(user.ID, user.Name, user.Gender, user.Age)
// 	return user, nil
// }

// // Update one user from the DB by its id
// func (db *SqlClient) UpdateUser(user types.User) error {
// 	// execute the sql statement
// 	comand, err := db.Prepare("update users set name=$2, gender=$3, age=$4 where id=$1")
// 	if err != nil {
// 		return fmt.Errorf("unable to update the user:" + err.Error())
// 	}
// 	defer comand.Close()
// 	comand.Exec(user.ID, user.Name, user.Gender, user.Age)
// 	return nil
// }

// // Delete one user from the DB by its id
// func (db *SqlClient) DeleteUser(id string) error {
// 	// execute the sql statement
// 	comand, err := db.Prepare("delete from users where id=$1")
// 	if err != nil {
// 		return fmt.Errorf("unable to delete the user:" + err.Error())
// 	}
// 	defer comand.Close()
// 	comand.Exec(id)
// 	return nil
// }
