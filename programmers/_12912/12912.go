// Package _12912
// https://programmers.co.kr/learn/courses/30/lessons/12912
package _12912

import "math"

func solution(a int, b int) int64 {
	var sum int64

	if a == b {
		return int64(a)
	}
	for i := math.Min(float64(a), float64(b)); i <= math.Max(float64(a), float64(b)); i++ {
		sum += int64(i)
	}

	return sum
}
