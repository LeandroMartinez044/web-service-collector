package collectorsrv

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/LeandroMartinez044/web-service-collector/internal/core/domain"
	"github.com/LeandroMartinez044/web-service-collector/internal/core/ports"
	"github.com/sirupsen/logrus"
)

type service struct {
	ytldRepo            ports.YoutubeRepository
	collectorRepository ports.WordRepository
}

func New(
	youtubeRepo ports.YoutubeRepository,
	collectorRepository ports.WordRepository) ports.CollectorService {
	return &service{youtubeRepo, collectorRepository}
}

// videoId: generate subtitle URL
func (srv *service) StoreSubtitlesByVideoId(videoId string) error {

	// Create a new instance of logrus logger
	logger := logrus.New()

	// Generate youtube's subtitles by videoId and store it
	err := srv.ytldRepo.GenerateSubtitlesFile(videoId)

	if err != nil {
		logger.Info(err)
		return fmt.Errorf(err.Error())
	}

	key := strings.Split(videoId, "/watch?v=")[1]

	// Get file that it will be to read it.
	file, err := srv.ytldRepo.GetFile(key)
	if err != nil {
		logger.Info(err)
		return fmt.Errorf("error obtener subitutlo %d", err)
	}

	// Creates slice with the words of the file.
	words := extractWords(file, videoId)
	defer file.Close()

	//Put word in DynamoDb
	//for _, word := range words {
	word := words[5]
	srv.collectorRepository.Put(word.Word, word.Sentence, word.Video.ID, word.Video.StartTime, word.Video.EndTime)
	//}

	// Removes file
	srv.ytldRepo.RemoveFile(key)

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
		if linea == "" || linea == "WEBVTT" || linea == "Kind: captions" || linea == "Language: en" || linea == "♪" || linea == "♪♪♪" {
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
		linea = strings.Replace(linea, "!", "", -1)

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
