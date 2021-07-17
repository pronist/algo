package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type MySignal struct {
	message string
}

func (s MySignal) String() string {
	return s.message
}

func (s MySignal) Signal() {}

//func doMain() {
//	defer fmt.Println("done infinite loop")
//	for {
//		time.Sleep(1 * time.Second)
//	}
//}

func main() {
	defer fmt.Println("Done")

	trapSignals := []os.Signal{
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, trapSignals...)

	//go doMain()

	time.AfterFunc(time.Second * 3, func() {
		sigCh <- MySignal{"Hello, Go!"}
	})

	sig := <-sigCh
	switch s := sig.(type) {
	case syscall.Signal:
		log.Println("Got signal", s)
	case MySignal:
		log.Println(s)
	}
}
