package main

import "fmt"

func main() {
	m := make(map[string]string)
	_, ok := m["Hello, Go!"]

	fmt.Println(ok)
}
