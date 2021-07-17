package main

import (
	"bytes"
	"fmt"
	"strings"
)

func solution(a, b int) bytes.Buffer {
	var buf bytes.Buffer

	for i := 0; i < b; i++ {
		fmt.Fprintln(&buf, strings.Repeat("*", a))
	}

	return buf
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	buf := solution(a, b)
	fmt.Print(buf.String())
}
