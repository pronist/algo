// Package _12933
// https://programmers.co.kr/learn/courses/30/lessons/12933
package _12933

import (
	"sort"
	"strconv"
	"strings"
)

func solution(n int64) int64 {
	r := strings.Split(strconv.FormatInt(n, 10), "")
	sort.Sort(sort.Reverse(sort.StringSlice(r)))

	answer, _ := strconv.Atoi(strings.Join(r, ""))
	return int64(answer)
}
