package main

import (
	"ch9/memo1"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() chan string {
	var r = make(chan string, 6)
	defer close(r)
	urls := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
	}
	for _, url := range urls {
		r <- url
	}
	return r
}

func main() {
	m := memo1.New(httpGetBody)
	var wg sync.WaitGroup
	for url := range incomingURLs() {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			log.Printf("%s, %s, %d bytes", url, time.Since(start), len(value.([]byte)))
			wg.Done()
		}(url)
	}
	wg.Wait()
}
