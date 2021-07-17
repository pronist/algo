// Package _12915
// https://programmers.co.kr/learn/courses/30/lessons/12915?language=go
package _12915

import (
	"sort"
)

type StringSlice struct {
	strings []string
	n       int
}

func (s StringSlice) Len() int {
	return len(s.strings)
}

func (s StringSlice) Less(i, j int) bool {
	if s.strings[i][s.n] == s.strings[j][s.n] {
		return s.strings[i] < s.strings[j]
	}
	return s.strings[i][s.n] < s.strings[j][s.n]
}

func (s StringSlice) Swap(i, j int) {
	s.strings[i], s.strings[j] = s.strings[j], s.strings[i]
}

func solution(strings []string, n int) []string {
	sort.Sort(StringSlice{strings, n})
	return strings
}