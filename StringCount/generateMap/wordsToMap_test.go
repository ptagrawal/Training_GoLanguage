package generateMap

import (
	"testing"
)

func TestGetMapFromWordsCaseInsensitive(t *testing.T) {
	countOfGO := 1
	GO := "GO"
	sentence := "This is a test Program for Go Language String Frequency Count"

	wordMap := GetMapFromWords(sentence)

	if wordMap[GO] != 1 {
		t.Error("Expected Count of word 'Go': %d, Got: %d", countOfGO, wordMap[GO])
	}
}

func TestGetMapFromWordsCaseSensitive(t *testing.T) {
	countOfGO := 4
	GO := "GO"
	sentence := "Go GO go gO"

	wordMap := GetMapFromWords(sentence)

	if wordMap[GO] != 4 {
		t.Error("Expected Count of word 'Go': %d, Got: %d", countOfGO, wordMap[GO])
	}
}

func TestGetMapFromWordsEmptyInput(t *testing.T) {
	countOfGO := 0
	GO := "GO"
	sentence := ""

	wordMap := GetMapFromWords(sentence)

	if wordMap[GO] != 0 {
		t.Error("Expected Count of word 'Go': %d, Got: %d", countOfGO, wordMap[GO])
	}
}
