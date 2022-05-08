// Go Class: 23 CSP, Goroutines, and Channels:
// https://youtu.be/zJd7Dvg3XCk?t=500
// very simialar to http://www.gopl.io/ch1.pdf , 1.6 Fetching URLs Concurrently

package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

func main() {
	results := make(chan result)
	list := []string{
		"https://amazon.com",
		"https://www.aliexpress.com",
		"https://srv-trade.ru",
		"https://uber4free.ru",
		"https://wsj.com",
	}

	for _, url := range list {
		go get(url, results)
	}

	for range list {
		r := <-results

		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err)
		} else {
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}
}
