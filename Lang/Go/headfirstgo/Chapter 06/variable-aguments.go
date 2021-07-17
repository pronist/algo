package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func main() {
	fmt.Println(maximum(71.8, 56.2, 32.1))

	numbers := []float64{71.8, 56.2, 32.1}
	fmt.Println(maximum(numbers...))
}
