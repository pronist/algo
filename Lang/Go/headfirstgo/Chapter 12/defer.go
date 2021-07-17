package main

import "fmt"

func Return() {
	defer fmt.Println("Return")
	return
}

func WithPanic() {
	defer fmt.Println("WithPanic")
	panic("Oops")
}

func main() {
	Return()
	WithPanic()
}
