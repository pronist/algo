// Package _12925
// https://programmers.co.kr/learn/courses/30/lessons/12925
package _12925

import "strconv"

func solution(s string) int {
	r, _ := strconv.ParseInt(s, 10, 64)
	return int(r)
}