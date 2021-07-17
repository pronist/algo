package main

import "fmt"

func main() {
    a := make([]int, 5)
    printSlice("a", a)

    // b := make([]int, 0, 5)
    b := []int{1, 2, 3, 4, 5}
    printSlice("b", b)

    c := b[:2]
    printSlice("c", c)

    d := c[2:5]
    printSlice("d", d)
}

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
