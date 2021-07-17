// Package
// https://programmers.co.kr/learn/courses/30/lessons/12934
package _12934

import (
	"math"
)

func solution(n int64) int64 {
	r := math.Sqrt(float64(n))

	if math.Mod(r, 1.0) == 0 {
		return int64(math.Pow(r+1, 2.0))
	}
	return -1
}
