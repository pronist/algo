package main

import (
	"flag"
	"fmt"
)

const (
	version = "1.0.0"
)

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()

	switch {
	case showVersion:
		fmt.Println(version)
	}
}
