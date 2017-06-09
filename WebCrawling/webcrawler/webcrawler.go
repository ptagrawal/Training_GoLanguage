package webcrawler

import (
	"fmt"
	"io"
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

func Crawl(pfunc func(Lang), lang Lang, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	err := getPage(&lang)
	if err != nil {
		fmt.Println("Error Received reading the page. %v", err)
	}
	elapsed := time.Since(start)
	lang.Time = elapsed.String()
	pfunc(lang)
}

func getPage(lang *Lang) error {
	resp, err := http.Get(lang.URL)
	if err != nil {
		return fmt.Errorf("Error Crawling the %v URL", lang.URL)
	}
	defer resp.Body.Close()

	count, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		return fmt.Errorf("Error Reading Bytes from the webpage.")
	}
	lang.Bytes = uint(count)

	return nil
}
