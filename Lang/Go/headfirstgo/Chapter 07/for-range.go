package main

import (
	"fmt"
	"sort"
)

func main() {
	grades := map[string]float64{"Alma": 74.2, "Rohit": 86.2, "Carl": 59.7}
	var names []string
	for name := range grades {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s has a grade of %.1f%%\n", name, grades[name])
	}
}
