package main

import (
	webcrawler "Training_GoLanguage/WebCrawling/webcrawler"
	"encoding/json"
	"fmt"
	"time"
)

var (
	Python = webcrawler.Lang{
		Name: "Python",
		URL:  "https://www.python.org/",
	}

	Ruby = webcrawler.Lang{
		Name: "Ruby",
		URL:  "https://www.ruby-lang.org/en/",
	}

	Golang = webcrawler.Lang{
		Name: "GoLang",
		URL:  "https://golang.org/",
	}
)

func main() {
	start := time.Now()
	webcrawler.Crawl(printLang, Python)
	webcrawler.Crawl(printLang, Ruby)
	webcrawler.Crawl(printLang, Golang)
	elapsed := time.Since(start)

	fmt.Printf("Total Time taken Crawling the three websites :" + elapsed.String() + "\n\n")
}

func printLang(lang webcrawler.Lang) {
	fmt.Printf("The Structure for %v in Go formatting is:\n%v\n\n", lang.Name, lang)
	outputJSON, err := json.Marshal(lang)
	if err != nil {
		fmt.Println("Received Error %v", err)
	}

	fmt.Printf("The Structure for %v in JSON formatting is:\n%v\n\n", lang.Name, string(outputJSON))
	fmt.Printf("\n")
}
