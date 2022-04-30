package models

import (
	"crudAPI/database"
	"crudAPI/entity"
	"log"

	"github.com/google/uuid"
)

// import (
// 	"GolangMongoDB/database"
// 	"GolangMongoDB/entity"
// 	"context"
// 	"fmt"
// 	"log"
// 	"strconv"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

// get one stock from the DB by its stockid
// func getAllStocks() ([]models.Stock, error) {
// 	// create the postgres db connection
// 	db := createConnection()

// 	// close the db connection
// 	defer db.Close()

// 	var stocks []models.Stock

// 	// create the select sql query
// 	sqlStatement := `SELECT * FROM stocks`

// 	// execute the sql statement
// 	rows, err := db.Query(sqlStatement)

// 	if err != nil {
// 		log.Fatalf("Unable to execute the query. %v", err)
// 	}

// 	// close the statement
// 	defer rows.Close()

// 	// iterate over the rows
// 	for rows.Next() {
// 		var stock models.Stock

// 		// unmarshal the row object to stock
// 		err = rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

// 		if err != nil {
// 			log.Fatalf("Unable to scan the row. %v", err)
// 		}

// 		// append the stock in the stocks slice
// 		stocks = append(stocks, stock)

// 	}

// 	// return empty stock on error
// 	return stocks, err
// }

// func GetUsers(ctx *gin.Context) (response *mongo.Cursor, err error) {
// 	recordPerPage, err := strconv.Atoi(ctx.Query("recordPerPage"))
// 	if err != nil || recordPerPage < 1 {
// 		recordPerPage = 10
// 	}

// 	startIndex, _ := strconv.Atoi(ctx.Query("startIndex"))
// 	matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}}
// 	groupStage := bson.D{{Key: "$group", Value: bson.D{
// 		{Key: "_id", Value: bson.D{{Key: "_id", Value: "null"}}},
// 		{Key: "total_count", Value: bson.D{{Key: "$sum", Value: 1}}},
// 		{Key: "data", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}}}}}
// 	projectStage := bson.D{
// 		{Key: "$project", Value: bson.D{
// 			{Key: "_id", Value: 0},
// 			{Key: "total_count", Value: 1},
// 			{Key: "user_items", Value: bson.D{
// 				{Key: "$slice", Value: []interface{}{"$data", startIndex, recordPerPage}}}}}}}

// 	var queryCtx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()
// 	response, err = userCollection.Aggregate(queryCtx, mongo.Pipeline{matchStage, groupStage, projectStage})
// 	if err != nil {
// 		log.Println(err.Error())
// 		return
// 	}
// 	return
// }

func CreateUser(user entity.User) error {
	db := database.PostgresInstance()
	defer db.Close()

	user.Id = uuid.New().String()
	comand, err := db.Prepare("insert into users(id,name,gender,age) values($1, $2, $3, $4)")
	if err != nil {
		log.Println("Unable to create user:", err.Error())
	}
	comand.Exec(user.Id, user.Name, user.Gender, user.Age)

	return nil
}

// func GetUser(UID string) (entity.User, error) {
// 	var user entity.User
// 	var queryCtx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Get a primitive ObjectID from a hexadecimal string
// 	idPrimitive, err := primitive.ObjectIDFromHex(UID)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return entity.User{}, err
// 	}

// 	// Call the FindOne() method by passing BSON
// 	if err := userCollection.FindOne(queryCtx, bson.M{"_id": idPrimitive}).Decode(&user); err != nil {
// 		return entity.User{}, err
// 	}
// 	return user, nil
// }

// func UpdateUser(id string, user entity.User) error {

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()
// 	// Declare a primitive ObjectID from a hexadecimal string
// 	idPrimitive, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return err
// 	}

// 	var updatedUser primitive.D
// 	updatedUser = append(updatedUser,
// 		bson.E{Key: "name", Value: user.Name},
// 		bson.E{Key: "gender", Value: user.Gender},
// 		bson.E{Key: "age", Value: user.Age})
// 	opt := options.Update().SetUpsert(true)
// 	update := bson.D{{Key: "$set", Value: updatedUser}}
// 	_, err = userCollection.UpdateByID(ctx, idPrimitive, update, opt)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return fmt.Errorf("unable to update user")
// 	}

// 	return nil
// }

// func DeleteUser(id string) error {
// 	var queryCtx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Declare a primitive ObjectID from a hexadecimal string
// 	idPrimitive, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return err
// 	}

// 	// Call the DeleteOne() method by passing BSON
// 	res, err := userCollection.DeleteOne(queryCtx, bson.M{"_id": idPrimitive})
// 	if err != nil {
// 		log.Println(err.Error())
// 		return fmt.Errorf("unable to delete user")
// 	} else if res.DeletedCount == 0 {
// 		return fmt.Errorf("there is no such user for be deleted")
// 	}
// 	return nil
// }

// //defer client.Disconnect(ctx)
