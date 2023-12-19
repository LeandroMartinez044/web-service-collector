package youtube

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// GenerateSubtitlesFile: saves subtitles in a file by VideoId
//
//   - videoId is the youtube's url of video,
//     Example: https://www.youtube.com/watch?v=Opxhh9Oh3rg. cannot be empty
//
// return error if something happens when trying to get the subtitles or try store it.
func (repo *repository) GenerateSubtitlesFile(videoId string) error {
	if videoId == "" {
		err := errors.New("videoId cannot be nil")
		return err
	}

	videoURL := videoId

	/*
		if err != nil {
			log.Fatal(err)
		}
	*/

	//Command to run youtube-dl to download subtitles
	cmd := exec.Command("youtube-dl", "--skip-download", "--write-sub", "--sub-lang", "en", "-o "+videoURL, videoURL)

	// Set output to os.Stdout to see the download progress
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Subtitles downloaded successfully.")
	return nil

}

// GetFileL: get file with subtitles by videoId.
//
//   - videoId is the youtube's url of video,
//     Example: https://www.youtube.com/watch?v=Opxhh9Oh3rg. cannot be empty
func (repo *repository) GetFile(videoId string) (*os.File, error) {
	if videoId == "" {
		return nil, errors.New("videoId cannot be empty")
	}

	fileLocation, err := repo.GetFileLocation(videoId)
	if err != nil {
		return nil, err
	}

	// Intenta abrir el archivo en modo lectura
	file, err := os.Open(fileLocation)
	if err != nil {
		// Maneja el error si no se pudo abrir el archivo
		return nil, err
	}

	return file, nil
}

// RemoveFile: remove download file by video id.
//
//   - videoId is the youtube's url of video,
//     Example: https://www.youtube.com/watch?v=Opxhh9Oh3rg. cannot be empty
//
// return nil if it remove successful otherwise if error it cannot remove the file.
func (repo *repository) RemoveFile(videoId string) error {

	fileLocation, err := repo.GetFileLocation(videoId)
	if err != nil {
		return err
	}

	err = os.Remove(fileLocation)
	if err != nil {
		return err
	}

	return nil
}

// GetFileLocation: builds subtitles file location with videoId.
//
//   - videoId is the youtube's url of video,
//     Example: https://www.youtube.com/watch?v=Opxhh9Oh3rg. cannot be empty
//
// return subtitles file location or an error if not receive the videoId.
func (repo *repository) GetFileLocation(videoId string) (string, error) {
	if videoId == "" {
		err := errors.New("videoId cannot be nil")
		return "", err
	}
	return repo.ytdl.Path + "/" + videoId + ".srt.en.vtt", nil
}

// buildDirectoryPath: builds the path for saving the subtitles.
//
// return directory path file. cannot be empty.
func buildDirectoryPath() string {
	directory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting currect directory:", err)
	}
	elementos := strings.Split(directory, "web-service-collector")
	return elementos[0] + "web-service-collector/resources"
}
