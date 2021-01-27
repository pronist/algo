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
	// divisor 로 나누어 떨어지는 element 가 하나도 없다면 배열에 -1 반환
	if len(result) == 0 {
		return []int{-1}
	}
	sort.Ints(result)

	return result
}
