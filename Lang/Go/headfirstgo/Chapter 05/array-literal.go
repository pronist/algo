package main

import "fmt"

func main() {
	arr := [4]int{0, 1, 2, 3}

	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])

	str := [4]byte{
		'A',
		'B',
		'C',
		'D',
	}
	fmt.Println(string(str[:]))
}
