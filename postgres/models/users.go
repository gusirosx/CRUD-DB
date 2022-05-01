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
	// create an empty user of type entity.User
	var user entity.User
	// create an empty list of type []entity.User
	var users []entity.User
	db := database.PostgresInstance()
	defer db.Close()
	// execute the sql statement
	rows, err := db.Query("select * from users")
	if err != nil {
		log.Println("Unable to execute the query: ", err.Error())
	}
	defer rows.Close()
	// iterate over the rows
	for rows.Next() {
		// unmarshal the row object to user
		err = rows.Scan(&user.Id, &user.Name, &user.Gender, &user.Age)
		if err != nil {
			// return empty users slice on error
			return []entity.User{}, fmt.Errorf("unable to retrieve the row:" + err.Error())
		}
		// append the user in the users slice
		users = append(users, user)
	}
	return users, nil
}

// Get one user from the DB by its id
func GetUser(UID string) (entity.User, error) {
	// create an empty user of type entity.User
	var user entity.User
	db := database.PostgresInstance()
	defer db.Close()
	// create the select sql query
	comand := "select * from users where id=$1"
	// execute the sql statement
	row := db.QueryRow(comand, UID)
	// unmarshal the row object to user struct
	if err := row.Scan(&user.Id, &user.Name, &user.Gender, &user.Age); err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// Create one user into DB
func CreateUser(user entity.User) error {
	db := database.PostgresInstance()
	defer db.Close()
	// get a unique userID
	user.Id = uuid.New().String()
	// execute the sql statement
	comand, err := db.Prepare("insert into users(id,name,gender,age) values($1, $2, $3, $4)")
	if err != nil {
		return fmt.Errorf("unable to create the user:" + err.Error())
	}
	defer comand.Close()
	comand.Exec(user.Id, user.Name, user.Gender, user.Age)
	return nil
}

// Update one user from the DB by its id
func UpdateUser(user entity.User) error {
	db := database.PostgresInstance()
	defer db.Close()
	// execute the sql statement
	comand, err := db.Prepare("update users set name=$2, gender=$3, age=$4 where id=$1")
	if err != nil {
		return fmt.Errorf("unable to update the user:" + err.Error())
	}
	defer comand.Close()
	comand.Exec(user.Id, user.Name, user.Gender, user.Age)
	return nil
}

// Delete one user from the DB by its id
func DeleteUser(id string) error {
	db := database.PostgresInstance()
	defer db.Close()
	// execute the sql statement
	comand, err := db.Prepare("delete from users where id=$1")
	if err != nil {
		return fmt.Errorf("unable to delete the user:" + err.Error())
	}
	defer comand.Close()
	comand.Exec(id)
	return nil
}
