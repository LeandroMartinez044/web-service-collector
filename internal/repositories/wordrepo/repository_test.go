package collector_test

import (
	"testing"

	"github.com/LeandroMartinez044/lmenglish/collector/internal/repositories/mongodb"
	wordrepo "github.com/LeandroMartinez044/lmenglish/collector/internal/repositories/wordrepo"
)

func TestGet(t *testing.T) {

	con := mongodb.New("mongodb://localhost:27017", "test")

	wordRepository := wordrepo.New(con)

	words, _ := wordRepository.Find("you")

	print(words)
}
