package main

import (
	"fmt"
	"reflect"
	"time"
)

func MakeTime(t reflect.Type) *time.Time {
	rv := reflect.New(t)
	return rv.Interface().(*time.Time)
}

func main() {
	t := MakeTime(reflect.TypeOf(time.Time{}))
	fmt.Println(t)
}
