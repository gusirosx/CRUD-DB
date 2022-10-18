package adapters

import (
	"clean2/domain/model"
	"clean2/usecase"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) usecase.UserRepository {
	return &userRepository{db}
}

// Get all users from the DB by its id
func (r *userRepository) FindAll() ([]*model.User, error) {
	// create an empty list of type []domain.User
	var users []*model.User

	// execute the sql statement
	rows, err := r.db.Query("select * from users")
	if err != nil {
		log.Println("Unable to execute the query: ", err.Error())
	}
	defer rows.Close()
	// iterate over the rows
	for rows.Next() {
		// create an empty user of type domain.User
		var user model.User
		// unmarshal the row object to user
		err = rows.Scan(&user.Id, &user.Name, &user.Gender, &user.Age)
		if err != nil {
			// return empty users slice on error
			return []*model.User{}, fmt.Errorf("unable to retrieve the row:" + err.Error())
		}
		// append the user in the users slice
		users = append(users, &user)
	}
	return users, nil
}

// Get one user from the DB by its id
func (r *userRepository) FindByID(UID string) (*model.User, error) {
	// create an empty user of type entity.User
	var user model.User
	// create the select sql query
	comand := "select * from users where id=$1"
	// execute the sql statement
	row := r.db.QueryRow(comand, UID)
	// unmarshal the row object to user struct
	if err := row.Scan(&user.Id, &user.Name, &user.Gender, &user.Age); err != nil {
		return &model.User{}, err
	}
	return &user, nil
}

// Get one user from the DB by its id
func (r *userRepository) Create(user *model.User) (*model.User, error) {
	// get a unique userID
	user.Id = uuid.New().String()
	// execute the sql statement
	comand, err := r.db.Prepare("insert into users(id,name,gender,age) values($1, $2, $3, $4)")
	if err != nil {
		return nil, fmt.Errorf("unable to create the user:" + err.Error())
	}
	defer comand.Close()
	comand.Exec(user.Id, user.Name, user.Gender, user.Age)

	return user, nil
}

// // Create one user into DB
// func CreateUser(user entity.User) error {
// 	// get a unique userID
// 	user.Id = uuid.New().String()
// 	// execute the sql statement
// 	comand, err := db.Prepare("insert into users(id,name,gender,age) values($1, $2, $3, $4)")
// 	if err != nil {
// 		return fmt.Errorf("unable to create the user:" + err.Error())
// 	}
// 	defer comand.Close()
// 	comand.Exec(user.Id, user.Name, user.Gender, user.Age)
// 	return nil
// }
