package main

import "fmt"

func main() {
	var mySlice []string
	fmt.Println(mySlice == nil)

	fmt.Println(len(mySlice))
	fmt.Println(append(mySlice, "Hello, Go!"))
}
