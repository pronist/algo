package main

import "fmt"

func main() {
	myArray := [3]string{"a", "b", "c"}
	mySlice := myArray[1:3]

	mySlice = append(mySlice, "d")
	fmt.Println(myArray, mySlice)
}
