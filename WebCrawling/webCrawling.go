package main

import (
	webcrawler "Training_GoLanguage/WebCrawling/webcrawler"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

var (
	langMap = make(map[string]webcrawler.Lang)
)

func main() {
	langMap["Python"] = webcrawler.Lang{
		Name: "Python",
		URL:  "https://www.python.org/",
	}

	langMap["Ruby"] = webcrawler.Lang{
		Name: "Ruby",
		URL:  "https://www.ruby-lang.org/en/",
	}

	langMap["Golang"] = webcrawler.Lang{
		Name: "GoLang",
		URL:  "https://golang.org/",
	}

	var wg sync.WaitGroup
	start := time.Now()
	for _, lang := range langMap {
		wg.Add(1)
		go webcrawler.Crawl(printLang, lang, &wg)
	}
	elapsed := time.Since(start)
	wg.Wait()

	fmt.Printf("Total Time taken Crawling the three websites : " + elapsed.String() + "\n\n")
}

func printLang(lang webcrawler.Lang) {
	fmt.Printf("The Structure for '%v' in Go formatting is:\n%v\n\n", lang.Name, lang)
	outputJSON, err := json.Marshal(lang)
	if err != nil {
		fmt.Println("Received Error %v", err)
	}

	fmt.Printf("The Structure for '%v' in JSON formatting is:\n%v\n\n", lang.Name, string(outputJSON))
	fmt.Printf("\n")
}
