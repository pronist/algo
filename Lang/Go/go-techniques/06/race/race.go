package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func main() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		mu.Lock()
		m["1"] = "a"
		mu.Unlock()
		c <- true
	}()
	mu.Lock()
	m["2"] = "b"
	mu.Unlock()

	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
