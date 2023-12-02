package collectorsrv

import (
	"testing"

	"github.com/LeandroMartinez044/lmenglish/collector/internal/repositories/youtube"
)

func TestGet(t *testing.T) {
	repo := youtube.New()
	srv := New(repo)

	//srv.ytldRepo.SaveSubtitules("https://www.youtube.com/watch?v=5NPBIwQyPWE")
	id := "https://www.youtube.com/watch?v=fFRl9sacyEQ"
	err := srv.StoreSubtitlesByVideoId(id)
	print(err)
}
