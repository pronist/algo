// Package _42758
// https://programmers.co.kr/learn/courses/30/lessons/42748
package _42748

import (
	"sort"
)

func solution(array []int, commands [][]int) []int {
	var result []int
	for _, cmd := range commands {
		p := append([]int{}, array[cmd[0]-1:cmd[1]]...)
		sort.Ints(p)

		result = append(result, p[cmd[2]-1])
	}
	return result
}
