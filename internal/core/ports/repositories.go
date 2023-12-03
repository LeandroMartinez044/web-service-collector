package ports

import (
	"os"

	"github.com/LeandroMartinez044/lmenglish/collector/internal/core/domain"
)

type YoutubeRepository interface {
	GenerateSubtitlesFile(videoUrl string) error
	GetFile(videoId string) (*os.File, error)
	RemoveFile(videoId string) error
	GetFileLocation(videoId string) (string, error)
}

type CollectorRepository interface {
	Save(word domain.Word)
	Find(word string) ([]domain.Word, error)
}
