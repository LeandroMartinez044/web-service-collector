package youtube

import (
	"github.com/LeandroMartinez044/web-service-collector/internal/core/ports"
	youtube_dl "github.com/youtube-videos/go-youtube-dl"
)

// - ytdl: represents a youtubeDl. Cannot be nil.
type repository struct {
	ytdl *youtube_dl.YoutubeDl
}

// New creates a new youtube repository.
//
// return YoutubeRepo. Cannot be nil.
func New() ports.YoutubeRepository {
	ytdl := &youtube_dl.YoutubeDl{}
	ytdl.Path = buildDirectoryPath()

	return &repository{ytdl: ytdl}
}
