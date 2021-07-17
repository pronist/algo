package main

import (
	"fmt"

	"example.com/pronist/geo"
)

func main() {
	location := geo.Coordinates{37.42, -122.08}
	fmt.Println("Latitude:", location.Latitude)
	fmt.Println("Longitude:", location.Longitude)
}
