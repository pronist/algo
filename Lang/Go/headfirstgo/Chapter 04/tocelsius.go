package main

import (
	"fmt"
	"log"

	"example.com/pronist/keyboard"
)

func main() {
	fmt.Print("Enter a temperature in Faherenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f degrees Celsius\n", celsius)
}
