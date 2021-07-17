package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type strSliceValue []string

func (v *strSliceValue) Set(s string) error {
	strs := strings.Split(s, ",")
	*v = append(*v, strs...)
	return nil
}

func (v *strSliceValue) String() string {
	return strings.Join(([]string)(*v), ",")
}

const (
	defaultPort = 3000
)

func main() {
	//var port int
	//flag.IntVar(&port, "port", defaultPort, "port to use")
	//flag.IntVar(&port, "p", defaultPort, "port to use (short)")
	//flag.Parse()

	flags := flag.NewFlagSet("example", flag.ContinueOnError)

	var species []string
	flags.Var((*strSliceValue)(&species), "species", "")

	if err := flags.Parse(os.Args[1:]); err != nil {
	}

	fmt.Printf("%#v", species)
	//flag.Parse()
}
