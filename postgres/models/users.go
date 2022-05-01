package models

import (
	"crudAPI/database"
	"crudAPI/entity"
	"fmt"
	"log"

	"github.com/google/uuid"
)

// Get all users from the DB by its id
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

// Get one user from the DB by its id
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

// Create one user into DB
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

// Update one user from the DB by its id
func UpdateUser(user entity.User) error {
	db := database.PostgresInstance()
	defer db.Close()
	// create the update sql query
	query := "update users set name=$2, gender=$3, age=$4 where id=$1"
	// execute the sql statement
	_, err := db.Exec(query, user.Id, user.Name, user.Gender, user.Age)
	if err != nil {
		return fmt.Errorf("unable to update the user :" + err.Error())
	}
	return nil
}

// Delete one user from the DB by its id
func DeleteUser(id string) error {
	db := database.PostgresInstance()
	defer db.Close()

	comand, err := db.Prepare("delete from users where id=$1")
	if err != nil {
		log.Println("Unable to delete user:", err.Error())
		return err
	}
	comand.Exec(id)
	return nil
}
