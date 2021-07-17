package main

import "fmt"

func main() {
	var pet struct {
		name string
		age  int
	}

	pet.name = "Max"
	pet.age = 5

	fmt.Printf("Name : %s\nAge %d", pet.name, pet.age)
}
