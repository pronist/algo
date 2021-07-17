package main

import (
	"fmt"
	"log"

	"example.com/pronist/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	votes := make(map[string]int, 2)

	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		votes[line]++
	}
	// fmt.Println(votes)

	for name, count := range votes {
		fmt.Printf("Name: %s, Counts: %d\n", name, count)
	}
}
