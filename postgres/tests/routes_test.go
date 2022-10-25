package tests

import (
	"bytes"
	"crudAPI/database"
	"crudAPI/handlers"
	"crudAPI/types"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

// Helper function to create a router during testing
func getRouterTesting() *gin.Engine {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

// Test the TestGetUsers handler function
func TestGetUsersHandler(t *testing.T) {
	// Get a new router
	r := getRouterTesting()
	// Define the route similar to its definition in the routes file
	r.GET("/users", handlers.GetUsers)
	// Create a response recorder
	req, _ := http.NewRequest("GET", "/users", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	// Test that the payload response is correct
	assert.Equal(t, http.StatusOK, res.Code)
}

// Test the TestGetUser handler function
func TestGetUserHandler(t *testing.T) {

	// add a user
	u := types.User{Name: "User Test", Gender: "robot", Age: 25}
	client := database.PostgresInstance()
	user, err := client.CreateUser(u)
	if err != nil {
		t.Fatalf("creation test failed: %v", err)
	}
	// remove the inserted users
	defer client.DeleteUser(u.ID)
	// Get a new router
	r := getRouterTesting()
	// Define the route similar to its definition in the routes file
	r.GET("/users/:id", handlers.GetUser)
	// Create a response recorder
	req, _ := http.NewRequest("GET", "/users/"+user.ID, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var usertest types.User
	_ = json.Unmarshal(res.Body.Bytes(), &usertest)
	// Test that the payload response is correct
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, usertest.ID, user.ID)

}

func TestCreateUserHandler(t *testing.T) {
	// Get a new router
	r := getRouterTesting()
	// Define the route similar to its definition in the routes file
	r.POST("/users", handlers.CreateUser)
	// add a user
	user := types.User{Name: "User Test", Gender: "robot", Age: 25}
	bytesUser, _ := json.Marshal(user)
	// Create a response recorder
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(bytesUser))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	// Test that the payload response is correct
	assert.Equal(t, http.StatusCreated, res.Code)
}

func TestDeleteUserHandler(t *testing.T) {
	// add a user
	u := types.User{Name: "User Test", Gender: "robot", Age: 25}
	client := database.PostgresInstance()
	user, err := client.CreateUser(u)
	if err != nil {
		t.Fatalf("creation test failed: %v", err)
	}
	// Get a new router
	r := getRouterTesting()
	// Define the route similar to its definition in the routes file
	r.DELETE("/users/:id", handlers.DeleteUser)
	// Create a response recorder
	req, _ := http.NewRequest("DELETE", "/users/"+user.ID, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	// Test that the payload response is correct
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestUpdateUserHandler(t *testing.T) {
	// add a user
	user := types.User{Name: "User Test", Gender: "robot", Age: 25}
	client := database.PostgresInstance()
	user, err := client.CreateUser(user)
	if err != nil {
		t.Fatalf("creation test failed: %v", err)
	}
	defer client.DeleteUser(user.ID)
	// modified user
	modUser := types.User{Name: "User Test", Gender: "robot", Age: 27}
	newBytesUser, _ := json.Marshal(modUser)
	// Get a new router
	r := getRouterTesting()
	// Define the route similar to its definition in the routes file
	r.PUT("/users/:id", handlers.UpdateUser)
	// Create a response recorder
	req, _ := http.NewRequest("PUT", "/users/"+user.ID, bytes.NewBuffer(newBytesUser))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	// Test that the payload response is correct
	assert.Equal(t, http.StatusOK, res.Code)
}
