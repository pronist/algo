package main

import "fmt"

func createPointer() *float64 {
	myFloat := 98.5
	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func main() {
	myFloatPointer := createPointer()
	fmt.Println(*myFloatPointer)

	var myBool bool
	printPointer(&myBool)
}
