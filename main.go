package main

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	os.Getenv("PING_URL")

	urls := []string{"https://www.linkedin.com/feed/", "https://www.google.com/"}

	for _, url := range urls {
		wg.Add(1)
		go pingUrl(url)
	}
	wg.Wait()
}

func pingUrl(url string) {

	for {
		_, err := http.Get(url)
		if err != nil {
			log.Println("Pinging Service Error", err)
			time.Sleep(5 * time.Second)
		}
		log.Println("Successful Pinging:", url)
	}

}
