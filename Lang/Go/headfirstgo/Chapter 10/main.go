package main

import (
	"fmt"

	"example.com/pronist/calendar"
)

func main() {
	event := calendar.Event{}

	event.SetYear(2019)
	event.SetDay(27)
	event.SetMonth(5)

	fmt.Println(event.Year())
}
