package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	X, Y int
}

func main() {
	p := &Point{10, 20}
	rv := reflect.ValueOf(p)

	fmt.Println(rv.Kind(), rv.Type(), rv.Interface())

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	xv := rv.Field(0)
	fmt.Println(xv.Int())
	xv.SetInt(100)
	fmt.Println(xv.Int())
}
