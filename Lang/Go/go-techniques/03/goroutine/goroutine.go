package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func fetchUrl(ctx context.Context, queue chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker exit")
			wg.Done()
			return
		case url := <-queue:
			fmt.Println("Start fetching URL", url)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	queue := make(chan string)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go fetchUrl(ctx, queue)
	}

	queue <- "http://example.com"
	queue <- "http://golang.org"
	queue <- "http://golang.org/doc"

	//close(queue)
	cancel()
	wg.Wait()
}
