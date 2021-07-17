package main

import (
	"fmt"

	"example.com/pronist/magazine"
)

func main() {
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.Street = "123 Oak St"
	subscriber.City = "Omaha"

	fmt.Println("Street:", subscriber.Street)
	fmt.Println("City:", subscriber.City)
}
