package main

import (
	"fmt"
	"reflect"
	"sort"
)

type Sorter struct {
	value reflect.Value
	fn func(int, int) bool
}

func (s Sorter) Len() int {
	return s.value.Len()
}

func (s Sorter) Less(i, j int) bool {
	return s.fn(i, j)
}

func (s Sorter) Swap(i, j int) {
	v1 := s.value.Index(i).Interface()
	v2 := s.value.Index(j).Interface()
	s.value.Index(j).Set(reflect.ValueOf(v1))
	s.value.Index(i).Set(reflect.ValueOf(v2))
}

func NewSorter(value interface{}, fn func(int, int) bool) *Sorter {
	return &Sorter {
		reflect.ValueOf(value),
		fn,
	}
}

func main() {
	numbers := []int{42, 10, 60, 20}
	s := NewSorter(numbers, func(i int, j int) bool {
		return numbers[i] < numbers[j]
	})
	sort.Sort(s)

	fmt.Println(numbers)
}