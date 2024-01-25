package domain

// word represents data about a word.
type Word struct {
	Word     string
	Sentence string
	Video    Video
}

func NewWord(word string, sentence string, video Video) *Word {
	return &Word{Word: word, Sentence: sentence, Video: video}
}
