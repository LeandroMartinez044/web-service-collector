package dependencies

import (
	"errors"
	"log"
	"os"

	"github.com/LeandroMartinez044/web-service-collector/internal/core/ports"
	"github.com/LeandroMartinez044/web-service-collector/internal/core/services/collectorsrv"
	"github.com/LeandroMartinez044/web-service-collector/internal/core/services/findersrv"
	"github.com/LeandroMartinez044/web-service-collector/internal/handlers/collectorhdl"
	wordrepo "github.com/LeandroMartinez044/web-service-collector/internal/repositories/wordrepo"
	"github.com/LeandroMartinez044/web-service-collector/internal/repositories/youtube"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Definition struct {

	//
	// Repositories
	//
	YoutubeRepository ports.YoutubeRepository
	WordRepository    ports.WordRepository

	//
	// Core
	//
	CollectorService ports.CollectorService
	FinderService    ports.FinderService

	//
	// Handlers
	//
	CollectorHandler *collectorhdl.Handler
}

func NewByEnvironment() Definition {

	svc, err := getDynamoClient()

	if err != nil {
		panic(errors.New("can't init application in development mode"))

	}

	return initDependencies(svc)
}

func initDependencies(svc *dynamodb.DynamoDB) Definition {

	d := Definition{}

	//
	// Repositories
	//
	d.WordRepository = wordrepo.New(svc, "words")
	d.YoutubeRepository = youtube.New()

	//
	// Core
	//
	d.CollectorService = collectorsrv.New(d.YoutubeRepository, d.WordRepository)
	d.FinderService = findersrv.New(d.WordRepository)

	//
	// Handlers
	//
	d.CollectorHandler = collectorhdl.New(d.CollectorService, d.FinderService)

	return d
}

func getDynamoClient() (*dynamodb.DynamoDB, error) {

	var sess *session.Session
	var err error
	region := "us-east-1"

	// Retrieve AWS region from environment variable or use a default value
	if os.Getenv("GO_ENV") == "test" {
		//
		dynamoDBEndpoint := "http://localhost:8080"
		// Create a new session using the AWS SDK for Go
		sess, err = session.NewSession(&aws.Config{
			Region:   aws.String(region),
			Endpoint: aws.String(dynamoDBEndpoint),
		})

	}

	sess, err = session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		log.Fatal("Error creating session:", err)
		return nil, err
	}

	/*
		print(os.Getenv("GO_ENV"))
		// Retrieve AWS region from environment variable or use a default value
		if os.Getenv("GO_ENV") == "prod" {
			// Create a new session using the AWS SDK for Go

		}
	*/
	if err != nil {
		return nil, err
	}

	// Create a DynamoDB client
	svc := dynamodb.New(sess)

	return svc, nil
}
