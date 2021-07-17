package main

import (
	"fmt"
	"reflect"
)

func main() {
	rv := reflect.ValueOf(map[string]int{"foo": 1})

	value := rv.MapIndex(reflect.ValueOf("foo"))
	fmt.Println(value.Int())

	rv.SetMapIndex(reflect.ValueOf("foo"), reflect.ValueOf(2))

	value = rv.MapIndex(reflect.ValueOf("foo"))
	fmt.Println(value.Int())
}