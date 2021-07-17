package main

import (
	"example.com/pronist/geo"
	"fmt"
	"log"
)

func main() {
	location := geo.Landmark{}
	err := location.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLongitude(-122.08)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(location.Latitude())
}
