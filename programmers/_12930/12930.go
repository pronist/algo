// Package _12930
// https://programmers.co.kr/learn/courses/30/lessons/12930
package _12930

import (
	"strings"
)

func solution(s string) string {
	var result []string
	for _, w := range strings.Split(s, " ") {
		var t string
		for i, r := range w {
			if i%2 == 0 {
				t += strings.ToUpper(string(r))
			} else {
				t += strings.ToLower(string(r))
			}
		}
		result = append(result, t)
	}

	return strings.Join(result, " ")
}
