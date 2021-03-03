// Package _12948
// https://programmers.co.kr/learn/courses/30/lessons/12948
package _12948

import "strings"

func solution(phone_number string) string {
	return strings.Repeat("*", len(phone_number) - 4) + phone_number[len(phone_number)-4:]
}