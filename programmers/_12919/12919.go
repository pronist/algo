// Package _12919
// https://programmers.co.kr/learn/courses/30/lessons/12919
package _12919

import (
	"strconv"
)

func solution(seoul []string) string {
	for i, s := range seoul {
		if s == "Kim" {
			return "김서방은 " + strconv.Itoa(i) + "에 있다"
		}
	}
	return ""
}
