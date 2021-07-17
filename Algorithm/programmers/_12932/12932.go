// Package _12932
// https://programmers.co.kr/learn/courses/30/lessons/12932
package _12932

import (
	"math"
	"strconv"
	"strings"
)

func solution(n int64) []int {
	var answer []int

	for _, c := range strings.Split(strconv.FormatInt(n, 10), "") {
		if t, err := strconv.Atoi(c); err == nil {
			answer = append(answer, t)
		}
	}
	for i, j := 0, len(answer)-1; i < int(math.Ceil(float64(len(answer)/2))); i, j = i+1, j-1 {
		answer[i], answer[j] = answer[j], answer[i]
	}
	return answer
}
