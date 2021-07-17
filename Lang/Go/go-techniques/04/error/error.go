package main

import (
	"fmt"
	"io"
	"os"
)

const (
	ExitCodeOK = 0
	ExitCodeError = 1
	ExitCodeFileError = 2
)

const (
	Version = "1.0"
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	fmt.Fprintf(c.outStream, "gobook version %s \n", Version)
	return 0
}

func main() {
	//// Not Executed
	//defer func() {
	//	fmt.Println("Hello, Go!")
	//}()
	//
	//os.Exit(ExitCodeError)

	cli := &CLI{os.Stdout, os.Stderr}
	os.Exit(cli.Run(os.Args))
}


