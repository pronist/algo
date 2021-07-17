// Package _12922
// https://programmers.co.kr/learn/courses/30/lessons/12922
package _12922

func solution(n int) string {
	var result string

	for i := 1; i <= n; i++ {
		if i % 2 != 0 {
			result += "수"
		} else {
			result += "박"
		}
	}
	return result
}