// Package _12928
// https://programmers.co.kr/learn/courses/30/lessons/12928
package _12928

func solution(n int) int {
	var sum int

	for i := 1; i <= n; i++ {
		if n % i == 0 {
			sum += i
		}
	}

	return sum
}