package collectorsrv_test

import (
	"testing"

	"github.com/LeandroMartinez044/web-service-collector/internal/core/services/collectorsrv"
	wordrepo "github.com/LeandroMartinez044/web-service-collector/internal/repositories/wordrepo"
	"github.com/LeandroMartinez044/web-service-collector/internal/repositories/youtube"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func TestGet(t *testing.T) {
	repo := youtube.New()
	endpoint := "http://localhost:8000" // Specify DynamoDB Local endpoint
	region := "us-east-1"               // Replace with your region

	sess, _ := session.NewSession(&aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(endpoint),
	})

	svc := dynamodb.New(sess)

	wordRepository := wordrepo.New(svc, "words")
	srv := collectorsrv.New(repo, wordRepository)

	//srv.ytldRepo.SaveSubtitules("https://www.youtube.com/watch?v=5NPBIwQyPWE")
	id := "https://www.youtube.com/watch?v=fFRl9sacyEQ"
	srv.StoreSubtitlesByVideoId(id)
}
