package youtube

import (
	"os"
	"testing"

	"github.com/LeandroMartinez044/web-service-collector/internal/core/ports"
)

var youtubeRepo ports.YoutubeRepository
var videoId string

func Setup() {
	youtubeRepo = New()
	videoId = "https://www.youtube.com/watch?v=Opxhh9Oh3rg"

}

func TearDown() {
	youtubeRepo.RemoveFile(videoId)
}

func TestGenerateSubtitlesFile(t *testing.T) {
	Setup()
	defer TearDown()

	err := youtubeRepo.GenerateSubtitlesFile(videoId)

	if err != nil {
		t.Errorf("SaveSubtitles(videoId) = %p", err)
	}

	fileLocation, _ := youtubeRepo.GetFileLocation(videoId)
	file, _ := os.Open(fileLocation)

	if file == nil {
		t.Errorf("file not generated")
	}

	defer file.Close()

}

func TestGetFile(t *testing.T) {
	Setup()
	defer TearDown()

	youtubeRepo.GenerateSubtitlesFile(videoId)

	file, _ := youtubeRepo.GetFile(videoId)

	if file == nil {
		t.Errorf("file not found")
	}

	youtubeRepo.RemoveFile(videoId)
}

func TestRemoveFile(t *testing.T) {
	Setup()
	youtubeRepo.GenerateSubtitlesFile(videoId)
	youtubeRepo.RemoveFile(videoId)

	file, _ := youtubeRepo.GetFile(videoId)

	if file != nil {
		t.Errorf("file not deleted")
	}
}
