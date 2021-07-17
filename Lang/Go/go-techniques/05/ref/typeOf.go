package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	X int `urlenc:"x,omitempty"`
	Y int `urlenc:"y,omitempty"`
}

func main() {
	rt := reflect.TypeOf(Point{})
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		if f.PkgPath != "" {}

		v, ok := f.Tag.Lookup("urlenc")
		if ok {
			fmt.Println(v)
		}
	}
}