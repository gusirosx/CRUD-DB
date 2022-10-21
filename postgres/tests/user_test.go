package tests

import (
	"crudAPI/database"
	"crudAPI/types"
	"testing"
)

func TestCreateUser(t *testing.T) {
	// userName := "teste 1"
	// userEmail := "teste1@mail.com"
	// userPassword := "test1-pass"
	// userPhone := "34 912341234"
	user := types.User{ID: "userHash", Name: "User Test", Gender: "robot", Age: 25}
	client := database.PostgresInstance()

	if err := client.CreateUser(user); err != nil {
		t.Fatalf("user creation failed: %v", err)
	}

	client.DeleteUser(user.ID)

	// t.Run("insert user on database", func(t *testing.T) {
	// 	CreateUser{
	// 	id, _ := InsertUser(userName, userEmail, userPassword, userPhone)
	// 	got := models.GetUser(id)
	// 	want := User{id, userName, userEmail, userPassword, userPhone}
	// 	assert.Equal(t, want, got)
	// 	models.DeleteUser(id)
	// })
	// t.Run("insert user with unique field already registered", func(t *testing.T) {
	// 	id, _ := InsertUser(userName, userEmail, userPassword, userPhone)
	// 	_, err := InsertUser(userName, userEmail, userPassword, userPhone)
	// 	assert.Error(t, err, "insert user with unique field already registered")
	// 	models.DeleteUser(id)
	// })
}

// func TestEditUserById(t *testing.T) {
// 	currentUser := User{
// 		Name:     "teste321",
// 		Email:    "teste321@mail",
// 		Password: "senha321",
// 		Phone:    "34 22222222",
// 	}
// 	modifiedUser := User{
// 		Name:     "teste123",
// 		Email:    "teste123@mail",
// 		Password: "senha123",
// 		Phone:    "34 phone",
// 	}

// 	t.Run("edit user", func(t *testing.T) {
// 		currentUser.ID, _ = InsertUser(
// 			currentUser.Name,
// 			currentUser.Email,
// 			currentUser.Password,
// 			currentUser.Phone,
// 		)
// 		modifiedUser.ID = currentUser.ID
// 		EditUser(modifiedUser)
// 		expectedUser := GetUserById(modifiedUser.ID)
// 		defer DeleteUserById(modifiedUser.ID)
// 		assert.Equal(t, expectedUser, modifiedUser)
// 	})
// }
