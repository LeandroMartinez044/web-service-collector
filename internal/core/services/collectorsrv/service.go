package collectorsrv

import (
	"bufio"
	"os"
	"strings"

	"github.com/LeandroMartinez044/lmenglish/collector/internal/core/domain"
	"github.com/LeandroMartinez044/lmenglish/collector/internal/core/ports"
)

type service struct {
	ytldRepo            ports.YoutubeRepository
	collectorRepository ports.CollectorRepository
}

func New(
	youtubeRepo ports.YoutubeRepository,
	collectorRepository ports.CollectorRepository) ports.CollectorService {
	return &service{youtubeRepo, collectorRepository}
}

// videoId: generate subtitle URL
func (srv *service) StoreSubtitlesByVideoId(videoId string) error {

	// Generate youtube's subtitles by videoId and store it
	err := srv.ytldRepo.GenerateSubtitlesFile(videoId)
	if err != nil {
		return err
	}

	// Get file that it will be to read it.
	file, err := srv.ytldRepo.GetFile(videoId)
	if err != nil {
		return err
	}

	// Creates slice with the words of the file.
	words := extractWords(file, videoId)
	defer file.Close()

	for _, word := range words {
		srv.collectorRepository.Save(word)
	}

	// Removes file
	srv.ytldRepo.RemoveFile(videoId)

	return nil
}

func extractWords(file *os.File, videoId string) []domain.Word {
	// Crea un nuevo lector bufio
	lector := bufio.NewScanner(file)

	// Lee el archivo línea por línea
	var words []domain.Word
	var video *domain.Video

	for lector.Scan() {
		linea := lector.Text()
		if linea == "" || linea == "WEBVTT" || linea == "Kind: captions" || linea == "Language: en" || linea == "♪" {
			continue
		}

		index := strings.Index(linea, "-->")
		if index != -1 {
			// Dividir la cadena en función de la coma "-->"
			elementos := strings.Split(linea, "-->")
			start := strings.TrimSpace(elementos[0])
			end := strings.TrimSpace(elementos[1])

			video = domain.NewVideo(videoId, start, end)
			continue
		}

		// Utilizar strings.Replace para eliminar el símbolo
		linea = strings.Replace(linea, "-", "", -1)

		// Utilizar strings.Replace para eliminar el símbolo
		linea = strings.Replace(linea, ".", "", -1)

		// Recorrer el slice utilizando la forma mejorada del bucle for
		for _, w := range strings.Fields(linea) {
			word := domain.NewWord(w, linea, *video)
			words = append(words, *word)
		}

	}

	return words
}
