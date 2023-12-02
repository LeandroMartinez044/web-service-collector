package domain

// word represents data about a word.
type Word struct {
	word  string
	video Video
}

func NewWord(word string, video Video) *Word {
	return &Word{word: word, video: video}
}
