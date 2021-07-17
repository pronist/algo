package main

import "fmt"

func sum(numbers ...int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 4))
	fmt.Println(sum(9, 7))
}
