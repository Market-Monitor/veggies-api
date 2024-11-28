package database

import (
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect() *mongo.Client {
	// Setup MongoDB client
	mongoClient, _ := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	return mongoClient
}
