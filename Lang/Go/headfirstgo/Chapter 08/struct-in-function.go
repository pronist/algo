package main

import "fmt"

type part struct {
	description string
	count       int
}

func showInfo(p part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)
}

func main() {
	p := part{description: "Hex bolts", count: 24}
	showInfo(p)
}
