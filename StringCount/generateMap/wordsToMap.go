package generateMap

import (
	"strings"
)

func GetMapFromWords(sentence string) map[string]int {
	// To make it case-independent
	sentence = strings.ToUpper(sentence)
	words := strings.Split(sentence, " ")

	wordMap := make(map[string]int)

	for _, word := range words {
		_, exists := wordMap[word]
		if !exists {
			wordMap[word] = 1
		} else {
			wordMap[word]++
		}
	}

	return wordMap
}
