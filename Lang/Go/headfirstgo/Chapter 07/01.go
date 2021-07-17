package main

import "fmt"

func main() {
	jewelry := make(map[string]float64)
	jewelry["necklace"] = 89.99
	jewelry["earrings"] = 79.99

	clothing := map[string]float64{"pants": 59.99, "shirt": 39.99}
	fmt.Printf("Earrings: %.2f\n", jewelry["earrings"])
	fmt.Printf("Necklace: %.2f\n", jewelry["necklace"])
	fmt.Printf("Shirt: %.2f\n", clothing["shirt"])
	fmt.Printf("Pants: %.2f\n", clothing["pants"])
}
