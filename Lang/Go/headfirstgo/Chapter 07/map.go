package main

import "fmt"

func main() {
	var ranks map[string]int
	ranks = make(map[string]int)

	ranks["gold"] = 1
	ranks["sliver"] = 2
	ranks["bronze"] = 3

	fmt.Println(ranks["bronze"])
	fmt.Println(ranks["gold"])
}
