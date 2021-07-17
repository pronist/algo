package main

import "fmt"

func main() {
	var notes []string
	notes = make([]string, 7)
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"

	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 3

	fmt.Println(len(notes))
	fmt.Println(len(primes))
}
