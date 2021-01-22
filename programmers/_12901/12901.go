// Package _12901
// https://programmers.co.kr/learn/courses/30/lessons/12901
package _12901

import (
	"strings"
	"time"
)

func solution(a int, b int) string {
	t := time.Date(2016, time.Month(a), b, 0, 0, 0, 0, time.UTC)
	return strings.ToUpper(t.Weekday().String()[:3])
}
