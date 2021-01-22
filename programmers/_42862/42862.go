// Package _42862
// https://programmers.co.kr/learn/courses/30/lessons/42862
package _42862

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

func solution(n int, lost []int, reserve []int) int {
	for p := 0; p < len(lost); p++ {
		for q := 0; q < len(reserve); q++ {
			if reserve[q] == lost[p] {
				lost = remove(lost, p)
				p--
				reserve = remove(reserve, q)
				q--
				break
			}
		}
	}

	can := n - len(lost)

	for p := 0; p < len(lost); p++ {
		for q := 0; q < len(reserve); q++ {
			if reserve[q] == lost[p]-1 || reserve[q] == lost[p]+1 {
				can++
				lost = remove(lost, p)
				p--
				reserve = remove(reserve, q)
				q--
				break
			}
		}
	}
	return can
}
