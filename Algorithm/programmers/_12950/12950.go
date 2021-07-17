// Package _12950
// https://programmers.co.kr/learn/courses/30/lessons/12950
package _12950

func solution(arr1 [][]int, arr2 [][]int) [][]int {
	answer := make([][]int, len(arr1))

	for i := 0; i < len(arr1); i++ {
		answer[i] = make([]int, len(arr1[i]))
		for j := 0; j < len(arr1[i]); j++ {
			answer[i][j] = arr1[i][j] + arr2[i][j]
		}
	}

	return answer
}
