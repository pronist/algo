package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("height can't be negative")
	fmt.Println(err.Error())
}
