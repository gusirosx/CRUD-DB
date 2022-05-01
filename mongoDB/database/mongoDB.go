package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create an exported global variable to hold the database connection pool.
var Client *mongo.Client = MongoInstance()

func MongoInstance() *mongo.Client {
	MongoDb := os.Getenv("MONGODB")
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatalln(err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Connected to MongoDB!")
	return client
}

func OpenCollection(client *mongo.Client, name string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("UsersMongo").Collection(name)
	return collection
}
