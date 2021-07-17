package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	X, Y int
}

func main() {
	rv := reflect.ValueOf(&Point{10, 20})

	if f := rv.Elem().Field(0); f.CanSet() {
		f.SetInt(100)
	}
	fmt.Println(rv.Interface().(*Point))
}
