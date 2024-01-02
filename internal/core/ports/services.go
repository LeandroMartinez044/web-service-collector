package ports

import "github.com/LeandroMartinez044/web-service-collector/internal/core/domain"

type CollectorService interface {
	StoreSubtitlesByVideoId(videoId string) error
}

type FinderService interface {
	FindVideosBy(word string) ([]domain.Video, error)
}
