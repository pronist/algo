package main

import "fmt"

var packageVar = "package"

func main() {
	var functionVar = "function"
	if true {
		var conditionVar = "conditinal"
		fmt.Println(packageVar)
		fmt.Println(functionVar)
		fmt.Println(conditionVar)
	}
	fmt.Println(packageVar)
	fmt.Println(functionVar)
	// fmt.Println(conditionVar)
}
