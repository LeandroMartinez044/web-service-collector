package collectorsrv_test

import (
	"testing"

	"github.com/LeandroMartinez044/lmenglish/collector/internal/core/services/collectorsrv"
	"github.com/LeandroMartinez044/lmenglish/collector/internal/repositories/collector"
	"github.com/LeandroMartinez044/lmenglish/collector/internal/repositories/mongodb"
	"github.com/LeandroMartinez044/lmenglish/collector/internal/repositories/youtube"
)

func TestGet(t *testing.T) {
	repo := youtube.New()
	con := mongodb.New("mongodb://localhost:27017", "test")

	collectorRepository := collector.New(con.Db.Collection("words"))
	srv := collectorsrv.New(repo, collectorRepository)

	//srv.ytldRepo.SaveSubtitules("https://www.youtube.com/watch?v=5NPBIwQyPWE")
	id := "https://www.youtube.com/watch?v=fFRl9sacyEQ"
	err := srv.StoreSubtitlesByVideoId(id)
	print(err)
}
