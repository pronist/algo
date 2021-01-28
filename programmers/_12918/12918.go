package _12918

import (
	"regexp"
)

func solution(s string) bool {
	if l := len(s); l == 4 || l == 6 {
		if matched, _ := regexp.MatchString(`^(\d)*$`, s); matched {
			return true
		}
	}
	return false
}
