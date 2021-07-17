// https://www.acmicpc.net/problem/1000
package main

import (
	"fmt"
	"log"
	"os"
)

func add(a, b int) int {
	return a + b
}

func main() {
	var a, b int

	_, err := fmt.Fscanf(os.Stdin, "%d %d", &a, &b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(add(a, b))
}
