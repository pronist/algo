// Package _12931
// https://programmers.co.kr/learn/courses/30/lessons/12931
package _12931

import (
	"strconv"
)

func solution(n int) int {
	var answer int

	for _, c := range strconv.Itoa(n) {
		s, _ := strconv.Atoi(string(c))
		answer += s
	}

	return answer
}
