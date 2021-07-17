package main

import "fmt"

type MyType string

func (m MyType) SayHi() {
	fmt.Println("Hi from", m)
}

func main() {
	value := MyType("a MyType value")
	value.SayHi()
}
