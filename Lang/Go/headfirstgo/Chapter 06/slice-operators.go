package main

import "fmt"

func main() {
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice1 := underlyingArray[1:4]

	fmt.Printf("%#v\n", slice1)

	slice1[0] = "A"
	fmt.Printf("%#v\n", underlyingArray)
}
