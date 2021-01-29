// Package _12926
// https://programmers.co.kr/learn/courses/30/lessons/12926
package _12926

import (
	"strings"
)

func solution(s string, n int) string {
	return strings.Map(func(r rune) rune {
		a := r + rune(n)
		switch {
		case r == ' ':
			return r
		case r >= 'a' && r <= 'z':
			if a > 'z' {
				return 'a' + (r - 'z' + rune(n) - 1)
			}
		case r >= 'A' && r <= 'Z':
			if a > 'Z' {
				return 'A' + (r - 'Z' + rune(n) - 1)
			}
		}
		return r + rune(n)
	}, s)
}