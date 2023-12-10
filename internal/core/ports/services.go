package ports

import "github.com/LeandroMartinez044/lmenglish/collector/internal/core/domain"

type CollectorService interface {
	StoreSubtitlesByVideoId(videoId string) error
}

type FinderService interface {
	FindVideosBy(word string) ([]domain.Video, error)
}
