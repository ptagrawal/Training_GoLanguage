package webcrawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Lang struct {
	Name  string
	URL   string
	Bytes uint
	Time  string
}

func Crawl(pfunc func(Lang), lang Lang, wg *sync.WaitGroup, dataCh chan byte) {
	defer wg.Done()
	start := time.Now()
	err := getPage(&lang, dataCh)
	if err != nil {
		fmt.Println("Error Received reading the page. %v", err)
	}
	elapsed := time.Since(start)
	lang.Time = elapsed.String()
	pfunc(lang)
}

func getPage(lang *Lang, dataCh chan byte) error {
	resp, err := http.Get(lang.URL)
	if err != nil {
		return fmt.Errorf("Error Crawling the %v URL", lang.URL)
	}
	defer resp.Body.Close()

	receivedBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error Reading the bytes from the webpage")
	}

	count := 0
	for _, bytes := range receivedBytes {
		dataCh <- bytes
		count++
	}

	// count, err := io.Copy(ioutil.Discard, resp.Body)
	// if err != nil {
	// 	return fmt.Errorf("Error Reading Bytes from the webpage.")
	// }

	lang.Bytes = uint(count)

	return nil
}
