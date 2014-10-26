package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var empty struct{}

func getStatus(urls []string) <-chan string {
	statusChan := make(chan string, 3)
	limit := make(chan struct{}, 5)
	go func() {
		for _, url := range urls {
			select {
			case limit <- empty:
				go func(url string) {
					res, err := http.Get(url)
					if err != nil {
						log.Fatal(err)
					}
					statusChan <- res.Status
					<-limit
				}(url)
			}
		}
	}()
	return statusChan
}

func main() {
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
		"http://denkinovel.com",
		"http://katryo.com",
	}
	time.Sleep(5 * time.Second)
	statusChan := getStatus(urls)
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}
}
