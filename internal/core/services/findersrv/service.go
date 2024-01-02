package findersrv

import (
	"github.com/LeandroMartinez044/web-service-collector/internal/core/domain"
	"github.com/LeandroMartinez044/web-service-collector/internal/core/ports"
)

type service struct {
	wordRepository ports.WordRepository
}

func New(wordRepository ports.WordRepository) service {
	return service{wordRepository: wordRepository}
}

type VideoDTO struct {
	Url       string `json:"url"`
	StartTime string `json:"start_time" `
	EndTime   string `json:"end_time"`
}

func (srv service) FindVideosBy(word string) ([]domain.Video, error) {
	words, err := srv.wordRepository.Find(word)

	var videos []domain.Video

	for i := 0; i < len(words); i++ {
		videos = append(videos, words[i].Video)
	}

	if err != nil {
		print(err)
	}

	return videos, nil
}
