package models

import (
	"crudAPI/database"
	"crudAPI/entity"
	"log"

	"github.com/google/uuid"
)

// get all users from the DB by its id
func GetUsers() ([]entity.User, error) {
	var user entity.User
	var users []entity.User
	db := database.PostgresInstance()
	defer db.Close()
	// execute the sql statement
	comand, err := db.Query("select * from users")
	if err != nil {
		log.Println("Unable to execute the query: ", err.Error())
	}
	// iterate over the rows
	for comand.Next() {
		// unmarshal the row object to user
		err = comand.Scan(&user.Id, &user.Name, &user.Gender, &user.Age)
		if err != nil {
			log.Println("Unable to scan the row: ", err.Error())
			// return empty users slice on error
			return []entity.User{}, nil
		}
		// append the user in the users slice
		users = append(users, user)
	}
	return users, nil
}

// get one user from the DB by its id
func GetUser(UID string) (entity.User, error) {
	var user entity.User
	db := database.PostgresInstance()
	defer db.Close()

	comand, err := db.Query("select * from users where id=$1", UID)
	if err != nil {
		log.Println("Error:", err.Error())
	}

	for comand.Next() {
		err = comand.Scan(&user.Id, &user.Name, &user.Gender, &user.Age)
		if err != nil {
			log.Println("Error:", err.Error())
		}
	}
	return user, nil
}

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
