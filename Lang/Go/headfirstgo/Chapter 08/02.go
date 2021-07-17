package main

import "fmt"

type student struct {
	name  string
	grade float64
}

func main() {
	s := student{name: "Alonzo Cole", grade: 92.3}
	fmt.Println("Name:", s.name)
	fmt.Println("Garde:", s.grade)
}
