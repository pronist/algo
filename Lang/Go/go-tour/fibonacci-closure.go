package main

import "fmt"

func fibonacci() func() int {
	var x, y int = 0, 1
	var z int

	return func() int {
		z, x, y = x, y, x+y
		return z
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
