package usecases

import (
	"CRUD-DB/clean/domain"
	"fmt"
	"log"
)

// Get all users from the DB by its id
func (r *UserRepoImpl) GetUsers() ([]*domain.User, error) {
	// create an empty user of type domain.User
	var user domain.User
	// create an empty list of type []domain.User
	var users []*domain.User

	// execute the sql statement
	rows, err := r.DB.Query("select * from users")
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
			return []*domain.User{}, fmt.Errorf("unable to retrieve the row:" + err.Error())
		}
		// append the user in the users slice
		users = append(users, &user)
	}
	return users, nil
}

// Get one user from the DB by its id
func (r *UserRepoImpl) GetUser(UID string) (*domain.User, error) {
	// create an empty user of type entity.User
	var user domain.User
	// create the select sql query
	comand := "select * from users where id=$1"
	// execute the sql statement
	row := r.DB.QueryRow(comand, UID)
	// unmarshal the row object to user struct
	if err := row.Scan(&user.Id, &user.Name, &user.Gender, &user.Age); err != nil {
		return &domain.User{}, err
	}
	return &user, nil
}
