package main

import "fmt"

func main() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}
	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true

	fmt.Printf("%#v", myStruct)
}
