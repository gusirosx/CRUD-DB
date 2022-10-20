package tests

import (
	"github.com/gin-gonic/gin"
)

// Helper function to create a router during testing
func getRouter() *gin.Engine {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

// Test the distance handler function
// func TestHandlerDistance(t *testing.T) {
// 	// Add some drivers
// 	client := database.PostgresInstance()

// 	// Get a new router
// 	r := getRouter()
// 	// Define the route similar to its definition in the routes file
// 	r.GET("/distance", handlers.Distance)

// 	req, err := http.NewRequest("GET", "/distance?"+getDistancePayload(), nil)
// 	if err != nil {
// 		t.Fatalf("could not create test request: %v", err)
// 	}
// 	// Create a response recorder
// 	w := httptest.NewRecorder()
// 	// Create the service and process the above request.
// 	r.ServeHTTP(w, req)

// 	// crate an anonymous struct for result response
// 	result := struct {
// 		Distance float64
// 	}{}

// 	// Test that the payload response is correct
// 	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
// 		t.Errorf("could not decode response %v", err)
// 	}

// 	if result.Distance != 333.6794 {
// 		t.Error("the first item in result could be driver 333.6794==========")
// 	}

// 	// Test that the http status code is 200
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	// Remove drivers
// 	client.RemovePoint("2")
// 	client.RemovePoint("3")
// }
