package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// numbers := [3]float64{71.8, 56.2, 89.5}
	// numbers, err := datafile.Getfloats("data.txt")

	numbers := os.Args[1:]
	var sum float64 = 0
	for _, number := range numbers {
		number, err := strconv.ParseFloat(number, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	cnt := float64(len(numbers))
	fmt.Printf("Average: %.2f\n", sum/cnt)
}
