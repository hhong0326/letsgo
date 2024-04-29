package main

import (
	"net/http"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	var urls = []string{
		"https://www.google.com",
		"https://www.naver.com",
		"https://www.daum.net",
	}

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			http.Get(url)
		}(url)
	}

	// All goroutines are done
	wg.Wait()
}
