package ports

import (
	"os"

	"github.com/LeandroMartinez044/web-service-collector/internal/core/domain"
)

type YoutubeRepository interface {
	GenerateSubtitlesFile(videoUrl string) error
	GetFile(videoId string) (*os.File, error)
	RemoveFile(videoId string) error
	GetFileLocation(videoId string) (string, error)
}

type WordRepository interface {
	Put(word string, sentence string, videoId string, videoStartTime string, videoEndTime string)
	Find(word string) ([]domain.Word, error)
}

type VideoRepository interface {
	Find(word string) ([]domain.Video, error)
}
