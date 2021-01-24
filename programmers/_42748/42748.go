// Package _42758
// https://programmers.co.kr/learn/courses/30/lessons/42748
package _42748

import (
	"sort"
)

func solution(array []int, commands [][]int) []int {
	var result []int

	for _, cmd := range commands {
		// p := array[cmd[0]-1:cmd[1]] 인 경우 같은 내부 공규 메모리 사용으로 정렬시 문제 발생
		p := append([]int{}, array[cmd[0]-1:cmd[1]]...)
		sort.Ints(p)

		result = append(result, p[cmd[2]-1])
	}
	return result
}
