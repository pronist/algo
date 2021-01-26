// Package _12917
// https://programmers.co.kr/learn/courses/30/lessons/12917
package _12917

import (
	"sort"
	"strings"
)

type StringSlice []string

func (p StringSlice) Len() int {
	return sort.StringSlice(p).Len()
}
func (p StringSlice) Less(i, j int) bool {
	return p[i] > p[j]
}
func (p StringSlice) Swap(i, j int) {
	sort.StringSlice(p).Swap(i, j)
}

func solution(s string) string {
	chars := StringSlice(strings.Split(s, ""))
	sort.Sort(chars)

	return strings.Join(chars, "")
}
