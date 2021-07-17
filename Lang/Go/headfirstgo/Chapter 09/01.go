package main

import "fmt"

type Population int

func main() {
	var population Population
	population = Population(572)
	fmt.Println("Sellpy Creek County population:", population)
	fmt.Println("Congratulations Kevin and Anna! It's a girl!")
	population += 1
	fmt.Println("Seepy Creek County population:", population)
}
