package ports

type CollectorService interface {
	StoreSubtitlesByVideoId(videoId string) error
}
