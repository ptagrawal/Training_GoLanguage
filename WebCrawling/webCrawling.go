package main

import (
	webcrawler "Training_GoLanguage/WebCrawling/webcrawler"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

func makeLangsList() []webcrawler.Lang {
	langs := make([]webcrawler.Lang, 0)

	langs = append(langs, webcrawler.Lang{
		Name: "Python",
		URL:  "https://www.python.org/",
	})

	langs = append(langs, webcrawler.Lang{
		Name: "Ruby",
		URL:  "https://www.ruby-lang.org/en/",
	})

	langs = append(langs, webcrawler.Lang{
		Name: "GoLang",
		URL:  "https://golang.org/",
	})

	return langs
}

func main() {
	var count int64
	langs := makeLangsList()

	dataChannel := make(chan byte)
	doneChannel := make(chan bool)

	go listenAndCount(dataChannel, doneChannel, &count)
	// go listenAndCount(dataChannel, &count)

	var wg sync.WaitGroup
	start := time.Now()
	for _, lang := range langs {
		wg.Add(1)
		go webcrawler.Crawl(printLang, lang, &wg, dataChannel)
	}
	wg.Wait()
	doneChannel <- true
	<-doneChannel
	elapsed := time.Since(start)

	fmt.Printf("Total number of bytes crawled from the %d websites : %d\n", len(langs), count)

	fmt.Printf("Total Time taken Crawling the three websites : " + elapsed.String() + "\n\n")
}

func listenAndCount(dataCh chan byte, doneCh chan bool, count *int64) {
	// func listenAndCount(dataCh chan byte, count *int64) {
	for {
		select {
		case <-doneCh:
			doneCh <- true
			return
		case <-dataCh:
			*count++
		}
	}
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
