// Package _12935
// https://programmers.co.kr/learn/courses/30/lessons/12935
package _12935

import "math"

func remove(s []int, i int) []int {
	switch {
	case len(s)-1 == i:
		return append(s[:len(s)-1])
	case i == 0:
		return append(s[1:])
	default:
		return append(s[:i], s[i+1:]...)
	}
}

func solution(arr []int) []int {
	if len(arr) == 1 {
		return []int{-1}
	}

	min := math.MaxInt64
	var minIndex int

	for i, num := range arr {
		if num < min {
			min, minIndex = num, i
		}
	}
	return remove(arr, minIndex)
}
