// Package _12954
// https://programmers.co.kr/learn/courses/30/lessons/12954
package _12954

func solution(x int, n int) []int64 {
	var answer []int64
	num := x

	for i := 0; i < n; num, i = num+x, i+1 {
		answer = append(answer, int64(num))
	}

	return answer
}
