package main

import (
	"bufio"
	"fmt"
	"github.com/mattn/go-isatty"
	"io"
	"os"
	"strings"
)

func main() {
	var output io.Writer
	if isatty.IsTerminal(os.Stdout.Fd()) {
		output = os.Stdout
	} else {
		output = bufio.NewWriter(os.Stdout)
	}
	for i := 0; i < 100; i++ {
		fmt.Fprintf(output, strings.Repeat("x", 100))
	}
	if _o, ok := output.(*bufio.Writer); ok {
		_o.Flush()
	}
}