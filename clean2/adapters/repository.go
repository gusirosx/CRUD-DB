package adapters

import (
	"clean2/domain/model"
	"clean2/usecase"
	"database/sql"
	"fmt"
	"log"
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
