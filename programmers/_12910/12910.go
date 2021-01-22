// Package _12910
// https://programmers.co.kr/learn/courses/30/lessons/12910
package _12910

import "sort"

func solution(arr []int, divisor int) []int {
	var result []int

	for _, v := range arr {
		if v%divisor == 0 {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return []int{-1}
	}
	sort.Ints(result)

	return result
}
