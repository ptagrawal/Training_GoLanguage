package main

import (
	getMap "Training_GoLanguage/StringCount/generateMap"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sentence string
	if len(os.Args) > 1 {
		args := os.Args[1:]
		sentence = strings.Join(args, " ")
	} else {
		input, err := getInputFromUser()
		if err != nil {
			fmt.Println("Received Error:", err)
		}
		sentence = input
	}

	wordMap := getMap.GetMapFromWords(sentence)
	generateOutput(wordMap)
}

func getInputFromUser() (string, error) {
	fmt.Println("Enter the Sentence or Series of Words separated with spaces")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("Error Reading the Input from the User")
	}
	inputWords := strings.TrimSpace(input)
	return inputWords, nil
}

func generateOutput(wordMap map[string]int) {
	fmt.Println("Output - \n[Word]:[Count]")
	for word, count := range wordMap {
		fmt.Println(word + " : " + strconv.Itoa(count))
	}
}
