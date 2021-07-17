package main

import "fmt"

func main() {
	m := map[string]int{"Hello, Go!": 10}
	fmt.Println(m)

	delete(m, "Hello, Go!")
	fmt.Println(m)
}
