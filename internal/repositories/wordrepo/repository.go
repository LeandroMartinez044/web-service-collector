package collector

import (
	"context"
	"fmt"
	"log"

	"github.com/LeandroMartinez044/lmenglish/collector/internal/core/domain"
	"github.com/LeandroMartinez044/lmenglish/collector/internal/repositories/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type repository struct {
	Connection *mongodb.Connection
}

func New(connection *mongodb.Connection) repository {
	return repository{Connection: connection}
}

func (r repository) Save(word domain.Word) {
	// Get a handle for the "users" collection in the "test" database

	// Create (Insert) Data
	insertResult, err := r.Connection.Db.Collection("words").InsertOne(context.Background(), word)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted ID: %v\n", insertResult.InsertedID)
}

func (r repository) Find(word string) ([]domain.Word, error) {

	// Find all documents where the "username" is "john_doe"
	filter := bson.M{"word": word}

	// Find documents matching the filter
	cursor, err := r.Connection.Db.Collection("words").Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	// Decode and print the results
	var words []domain.Word
	if err := cursor.All(context.Background(), &words); err != nil {
		log.Fatal(err)
	}

	return words, nil
}
