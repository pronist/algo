// Package _12940
// https://programmers.co.kr/learn/courses/30/lessons/12940
package _12940

import "math"

func solution(n int, m int) []int {
	a, b, gcd := n, m, 1

	for i := 2; i < int(math.Min(float64(a), float64(b))); i++ {
		if a % i == 0 && b % i == 0 {
			a, b, gcd, i = a/i, b/i, gcd*i, 1
		}
	}

	return []int{gcd, gcd * a * b}
}