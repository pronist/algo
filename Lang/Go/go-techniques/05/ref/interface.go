package main

import (
	"fmt"
	"reflect"
)

type MyInterface interface {
	Foo()
}

type MyType struct {}

func (t MyType) Foo() {
	fmt.Println("Hello, Go!")
}

func main() {
	rv := reflect.ValueOf(MyType{})

	var i MyInterface
	iv := reflect.TypeOf(&i).Elem()

	if rv.Type().Implements(iv) {
		rv.Interface().(MyInterface).Foo()
	}
}
