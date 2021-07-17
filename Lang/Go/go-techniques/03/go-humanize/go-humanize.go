package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"os"
)

func main() {
	name := os.Args[1]
	s, _ := os.Stat(name)

	fmt.Printf("%s: %s\n", name, humanize.Bytes(uint64(s.Size())))
}