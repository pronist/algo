package main

import "fmt"

func ExampleHello() {
	fmt.Println("Hello")
	// Output: Hello
}

func ExampleUnordered() {
	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
	}
	// Unordered output:
	// 1
	// 2
	// 3
}