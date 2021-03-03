// Package _12947
// https://programmers.co.kr/learn/courses/30/lessons/12947
package _12947

import "strconv"

func solution(x int) bool {
	var sum int

	for _, n := range strconv.FormatInt(int64(x), 10) {
		t, _ := strconv.Atoi(string(n))
		sum += t
	}

	return x % sum == 0
}