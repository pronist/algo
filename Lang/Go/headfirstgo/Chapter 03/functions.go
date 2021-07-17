package main

import "fmt"

func sayHi() {
	fmt.Println("Hi!")
}

func repeatLine(line string, time int) {
	for i := 0; i < time; i++ {
		fmt.Println(line)
	}
}

func paintNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)
}

func main() {
	// sayHi()
	// repeatLine("Hello", 3)

	paintNeeded(4.2, 3.0)
	paintNeeded(5.2, 3.5)
}
