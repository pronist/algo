package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := map[string]int{"a": 0, "b": 1}
	rv := reflect.ValueOf(m)

	for _, key := range rv.MapKeys() {
		v := rv.MapIndex(key)
		switch v.Kind() {
		case reflect.Int:
			fmt.Println(v.Int())
		}
	}
}
