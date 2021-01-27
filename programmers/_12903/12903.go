// Package _12903
// https://programmers.co.kr/learn/courses/30/lessons/12903
package _12903

import (
	"math"
)

func solution(s string) string {
	center := int(math.Floor(float64(len(s))) / 2)

	if len(s)%2 != 0 {
		return string(s[center])
	}
	return s[center-1 : center+1]
}