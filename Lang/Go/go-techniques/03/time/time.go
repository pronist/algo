package main

import (
	"fmt"
	"log"
	"time"
)

func doSomeThing() error {
	time.Sleep(time.Second * 2)
	return nil
}

func main() {
	timer := time.NewTimer(time.Second * 3)
	done := make(chan error)
	go func() {
		err := doSomeThing()
		done <- err
	}()
	select {
	case <-timer.C:
		fmt.Println("Hello, Go!")
	case err := <-done:
		if err != nil {
			log.Fatal(err)
		}
	}
}
