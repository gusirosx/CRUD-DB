package tests

import (
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
func TestGetUsers(t *testing.T) {
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
func TestGetUser(t *testing.T) {

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

// func TestDeleteUserById(t *testing.T) {
// 	r := setupTestingRoutes()
// 	r.DELETE("/users/:id", handlers.DeleteUserById)

// 	t.Run("test invalid id", func(t *testing.T) {
// 		req, _ := http.NewRequest("DELETE", "/users/asdf", nil)
// 		res := httptest.NewRecorder()

// 		r.ServeHTTP(res, req)
// 		expect := `{"error":"cannot understand this id"}`
// 		resBody, _ := ioutil.ReadAll(res.Body)
// 		got := string(resBody)
// 		assert.Equal(t, expect, got, "expected to got an error response body")
// 	})

// 	t.Run("delete not existent id", func(t *testing.T) {
// 		req, _ := http.NewRequest("DELETE", "/users/1", nil)
// 		res := httptest.NewRecorder()
// 		r.ServeHTTP(res, req)

// 		expect := `{"status":"user not found to delete"}`
// 		resBody, _ := ioutil.ReadAll(res.Body)
// 		got := string(resBody)
// 		assert.Equal(t, expect, got, "expected to got an not found response")
// 	})

// 	t.Run("delete existing id", func(t *testing.T) {
// 		userExpect := models.User{
// 			Name:     "test get by id",
// 			Email:    "getbyid@mail.com",
// 			Password: "pass",
// 			Phone:    "a number",
// 		}
// 		id, _ := models.InsertUser(
// 			userExpect.Name,
// 			userExpect.Email,
// 			userExpect.Password,
// 			userExpect.Phone,
// 		)
// 		userExpect.ID = id

// 		req, _ := http.NewRequest("DELETE", "/users/"+strconv.Itoa(id), nil)
// 		res := httptest.NewRecorder()
// 		r.ServeHTTP(res, req)

// 		userJson, _ := json.Marshal(userExpect)
// 		expect := "{\"data\":" + string(userJson) + ",\"status\":\"user deleted successfully\"}"
// 		resBody, _ := ioutil.ReadAll(res.Body)
// 		got := string(resBody)
// 		assert.Equal(t, expect, got, "expect a success message")
// 	})
// }

// func TestInsertUser(t *testing.T) {
// 	r := setupTestingRoutes()
// 	r.POST("/users", handlers.CreateUser)

// 	t.Run("create valid user", func(t *testing.T) {
// 		user := models.User{
// 			Name:     "test get by id",
// 			Email:    "getbyid@mail.com",
// 			Password: "pass",
// 			Phone:    "a number",
// 		}

// 		bytesUser, _ := json.Marshal(user)
// 		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(bytesUser))
// 		res := httptest.NewRecorder()
// 		r.ServeHTTP(res, req)
// 		resBody, _ := ioutil.ReadAll(res.Body)
// 		var userGot models.User
// 		_ = json.Unmarshal(resBody, &userGot)

// 		assert.Equal(t, http.StatusCreated, res.Code)
// 		assert.Equal(t, user.Email, userGot.Email)

// 		defer models.DeleteUserById(userGot.ID)

// 	})
// }

// func TestEditUserById(t *testing.T) {
// 	r := setupTestingRoutes()
// 	r.PUT("/users/:id", handlers.EditUserById)

// 	currentUser := models.User{
// 		Name:     "teste321",
// 		Email:    "teste321@mail",
// 		Password: "senha321",
// 		Phone:    "34 22222222",
// 	}
// 	modifiedUser := models.User{
// 		Name:     "teste123",
// 		Email:    "teste123@mail",
// 		Password: "senha123",
// 		Phone:    "34 phone",
// 	}

// 	t.Run("test update existing user", func(t *testing.T) {
// 		currentUser.ID, _ = models.InsertUser(
// 			currentUser.Name,
// 			currentUser.Email,
// 			currentUser.Password,
// 			currentUser.Phone,
// 		)
// 		defer models.DeleteUserById(currentUser.ID)

// 		modifiedUser.ID = currentUser.ID
// 		newBytesUser, _ := json.Marshal(modifiedUser)

// 		req, _ := http.NewRequest("PUT", "/users/"+strconv.Itoa(currentUser.ID), bytes.NewBuffer(newBytesUser))
// 		res := httptest.NewRecorder()
// 		r.ServeHTTP(res, req)

// 		resBody, _ := ioutil.ReadAll(res.Body)
// 		var userGot models.User
// 		_ = json.Unmarshal(resBody, &userGot)

// 		assert.Equal(t, http.StatusOK, res.Code, "status code dont match")
// 		assert.Equal(t, modifiedUser, userGot, "json body dont match")
// 	})

// 	t.Run("edit inexistent user id on database", func(t *testing.T) {
// 		modifiedUser.ID = 0

// 		newBytesUser, _ := json.Marshal(modifiedUser)

// 		req, _ := http.NewRequest("PUT", "/users/"+strconv.Itoa(0), bytes.NewBuffer(newBytesUser))
// 		res := httptest.NewRecorder()
// 		r.ServeHTTP(res, req)

// 		assert.Equal(t, http.StatusNotFound, res.Code)
// 	})
// }
