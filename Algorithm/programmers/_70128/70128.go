// Package _70128
// https://programmers.co.kr/learn/courses/30/lessons/70128
package _70128

func solution(a []int, b []int) int {
	var sum int

	for i := range a {
		sum += a[i] * b[i]
	}

	return sum
}
