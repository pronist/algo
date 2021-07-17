package main

import "fmt"

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	var s subscriber
	applyDiscount(&s)
	fmt.Println(s.rate)
}
