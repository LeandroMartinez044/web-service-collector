package collector

import (
	"fmt"
	"log"

	"github.com/LeandroMartinez044/web-service-collector/internal/core/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type repository struct {
	svc       *dynamodb.DynamoDB
	nameTable string
}

func New(dynamodb *dynamodb.DynamoDB, nameTable string) *repository {
	return &repository{svc: dynamodb, nameTable: nameTable}
}

func (r *repository) Put(word string, sentence string, videoId string,
	videoStartTime string, videoEndTime string) {

	// Define the item to put
	item := map[string]*dynamodb.AttributeValue{
		"word": {
			S: aws.String(word),
		},
		"sentence": {
			S: aws.String(sentence),
		},
		"video": {
			M: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String(videoId),
				},
				"starttime": {
					S: aws.String(videoStartTime),
				},
				"endtime": {
					S: aws.String(videoEndTime),
				},
			},
		},
	}

	// Put in the table words a word
	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.nameTable),
		Item:      item,
	}

	_, err := r.svc.PutItem(input)

	if err != nil {
		log.Fatal("Error saving item in DynamoDB:", err)
		return
	}

	fmt.Println("Item saved successfully!")

}

func (r repository) Find(word string) ([]domain.Word, error) {

	return nil, nil
}
