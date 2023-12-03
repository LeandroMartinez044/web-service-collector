package collector

import (
	"context"
	"fmt"
	"log"

	"github.com/LeandroMartinez044/lmenglish/collector/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	collection *mongo.Collection
}

func New(collection *mongo.Collection) repository {
	return repository{collection: collection}
}

func (r repository) Save(word domain.Word) {
	// Get a handle for the "users" collection in the "test" database

	// Create (Insert) Data
	insertResult, err := r.collection.InsertOne(context.Background(), word)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted ID: %v\n", insertResult.InsertedID)
}

func (r repository) Find(word string) ([]domain.Word, error) {

	// Find all documents where the "username" is "john_doe"
	filter := bson.M{"word": word}

	// Find documents matching the filter
	cursor, err := r.collection.Find(context.Background(), filter)
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
