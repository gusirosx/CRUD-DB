package services

import (
	"crudAPI/database"
	"crudAPI/types"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create an unexported global variable to hold the database connection pool.
var client *database.MongoClient = database.MongoInstance()

// Get one user from the DB by its id
func GetUsers(ctx *gin.Context) (primitive.M, error) {
	userList, err := client.GetUsers(ctx)
	if err != nil {
		return primitive.M{}, err
	}
	return userList, nil
}

// Get one user from the DB by its id
func GetUser(UID string) (types.User, error) {
	user, err := client.GetUser(UID)
	if err != nil {
		return types.User{}, err
	}
	return user, nil
}

// CreateUser create a user in the postgres database
func CreateUser(user types.User) (types.User, error) {
	user, err := client.CreateUser(user)
	if err != nil {
		return types.User{}, err
	}
	return user, nil
}

// UpdateUser updates a user in the postgres database
func UpdateUser(userID string, user types.User) error {
	if err := client.UpdateUser(userID, user); err != nil {
		return err
	}
	return nil
}

// // DeleteUser delete a user in the postgres database
// func DeleteUser(ID string) error {
// 	if err := client.DeleteUser(ID); err != nil {
// 		return err
// 	}
// 	return nil
// }
