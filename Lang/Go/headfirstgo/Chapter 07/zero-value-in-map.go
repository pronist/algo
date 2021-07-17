package main

import "fmt"

func main() {
	numbers := make(map[string]int)
	numbers["I've been assigned"] = 12

	fmt.Println(numbers["I've been assigned"])
	fmt.Println(numbers["I haven't been assigned"])

	var nilMap map[int]string
	fmt.Println(nilMap == nil)
}
