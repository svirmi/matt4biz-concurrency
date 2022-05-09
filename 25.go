// Go Class: 25 Context
// https://youtu.be/0x_oUlxzw5A?t=575

package main

import (
	"context"
	"log"
	"net/http"
	"runtime"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(ctx context.Context, url string, ch chan<- result) {
	start := time.Now()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if resp, err := http.DefaultClient.Do(req); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

func main() {

	for i := 0; i < 10; i++ {

		ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)

		defer cancel()

		list := []string{
			"https://amazon.com",
			"https://www.aliexpress.com",
			"https://vshkole5.ru",
			"https://uber4free.ru",
			"https://wsj.com",
			"https://1234wsj.com",
		}

		results := make(chan result, len(list))

		defer close(results)

		for _, url := range list {
			go get(ctx, url, results)
		}

		for range list {
			r := <-results

			if r.err != nil {
				log.Printf("%-20s %s\n", r.url, r.err)
			} else {
				log.Printf("%-20s %s\n", r.url, r.latency)
			}
		}

		log.Println(i)
	}

	log.Println(runtime.NumGoroutine(), " gouroutines still running")
}
