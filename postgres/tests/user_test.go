package tests

import (
	"crudAPI/database"
	"crudAPI/types"
	"testing"
)

func TestCreateUser(t *testing.T) {
	user := types.User{Name: "User Test", Gender: "robot", Age: 25}
	client := database.PostgresInstance()

	user, err := client.CreateUser(user)
	if err != nil {
		t.Fatalf("creation test failed: %v", err)
	}

	client.DeleteUser(user.ID)
}

func TestUpdateUser(t *testing.T) {
	user := types.User{Name: "User Test", Gender: "robot", Age: 25}
	client := database.PostgresInstance()

	user, err := client.CreateUser(user)
	if err != nil {
		t.Fatalf("creation test failed: %v", err)
	}

	modifiedUser := types.User{ID: "userHash", Name: "User Test modified", Gender: "robot", Age: 26}
	if err := client.UpdateUser(modifiedUser); err != nil {
		t.Fatalf("update test failed: %v", err)
	}

	client.DeleteUser(user.ID)
}
