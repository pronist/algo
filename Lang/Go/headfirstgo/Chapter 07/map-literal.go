package main

import "fmt"

func main() {
	ranks := map[string]int{"bronze": 3, "sliver": 2, "gold": 1}
	fmt.Println(ranks["gold"])
	fmt.Println(ranks["bronze"])

	elements := map[string]string{
		"H":  "Hydrogen",
		"Li": "Lithuum",
	}
	fmt.Println(elements["H"])
	fmt.Println(elements["Li"])
}
