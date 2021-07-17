package main

import (
	"fmt"
	"log"

	"example.com/pronist/datafile"
)

func main() {
	// numbers := [3]float64{71.8, 56.2, 89.5}

	numbers, err := datafile.Getfloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	cnt := float64(len(numbers))
	fmt.Printf("Average: %.2f\n", sum/cnt)
}
