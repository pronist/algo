package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

func ReadFromFile(ch chan []byte, f *os.File)  {
	defer close(ch)
	defer f.Close()

	 buf := make([]byte, 4096)
	 for {
	 	if _, err := f.Read(buf); err == nil {
	 		ch <- buf[:]
		}
	 }
}

func MakeChannelsFormFiles(files []string) ([]reflect.Value, error){
	cs := make([]reflect.Value, len(files))

	for i, fn := range files {
		ch := make(chan []byte)

		f, err := os.Open(fn)
		if err != nil {
			return nil, err
		}
		go ReadFromFile(ch, f)

		cs[i] = reflect.ValueOf(ch)
	}
	return cs, nil
}

func MakeSelectCases(cs ...reflect.Value) ([]reflect.SelectCase, error) {
	c := make([]reflect.SelectCase, len(cs))

	for i, ch := range cs {
		if ch.Kind() != reflect.Chan {
			return nil, errors.New("argument must be a channel")
		}
		c[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: ch }
	}
	return c, nil
}

func DoSelect(cases []reflect.SelectCase) {
	for {
		if chosen, recv, ok := reflect.Select(cases); ok {
			fmt.Printf("\n===%s===\n%s", os.Args[chosen+1], recv.Interface())
		}
	}
}

func main() {
	if err := _main(); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

func _main() error {
	if len(os.Args) < 2 {
		return errors.New("prog [file1 file2 ...]")
	}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	channels, err := MakeChannelsFormFiles(os.Args[1:])
	if err != nil {
		return err
	}

	cases, err := MakeSelectCases(channels...)
	if err != nil {
		return err
	}
	go DoSelect(cases)

	select {
	case <-sigch:
		return nil
	}
	return nil
}
