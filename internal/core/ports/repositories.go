package ports

import "os"

type YoutubeRepository interface {
	GenerateSubtitlesFile(videoUrl string) error
	GetFile(videoId string) (*os.File, error)
	RemoveFile(videoId string) error
	GetFileLocation(videoId string) (string, error)
}
