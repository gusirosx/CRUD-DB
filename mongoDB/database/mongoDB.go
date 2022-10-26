package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDb = os.Getenv("MONGODB")

// create the mongoDB database connection
func MongoInstance() *MongoClient {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatalln("Unable to establish connection:", err.Error())
	}
	log.Println("Connected to MongoDB!")
	return &MongoClient{client}
}

func (client *MongoClient) OpenCollection(name string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("UsersMongo").Collection(name)
	return collection
}

// Create a custom SqlClient type which wraps the redis.Client connection pool.
type MongoClient struct {
	*mongo.Client
}
