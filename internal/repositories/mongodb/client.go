package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	client *mongo.Client
	Db     *mongo.Database
}

func New(URI string, dbName string) *Connection {
	// Set client options
	clientOptions := options.Client().ApplyURI(URI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(dbName)

	return &Connection{client, database}
}
