package main

import "sort"

//type IntSlice []int
//
//func (p IntSlice) Len() int {
//	return len(p)
//}
//
//func (p IntSlice) Less(i, j int) bool {
//	return p[i] < p[j]
//}
//
//func (p IntSlice) Swap(i, j int) {
//	p[i], p[j] = p[j], p[i]
//}

func main() {
	numbers := []int{3, 4, 2, 6, 10}
	sort.Sort(sort.IntSlice(numbers))
}