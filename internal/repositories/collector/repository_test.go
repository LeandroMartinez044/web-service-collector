package collector_test

import (
	"testing"

	"github.com/LeandroMartinez044/lmenglish/collector/internal/repositories/collector"
	"github.com/LeandroMartinez044/lmenglish/collector/internal/repositories/mongodb"
)

func TestGet(t *testing.T) {

	con := mongodb.New("mongodb://localhost:27017", "test")

	collectorRepository := collector.New(con.Db.Collection("words"))

	words, _ := collectorRepository.Find("you")

	print(words)
}
