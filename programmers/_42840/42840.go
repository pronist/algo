// Package _42840
// https://programmers.co.kr/learn/courses/30/lessons/42840
package _42840

import (
	"math"
)

type p struct {
	pattern []int // 각 수포자의 패턴
	i       int // 내부 이터레이터
	count   int // 문제를 맞춘 횟수
}

func (p *p) resolve(answer int) {
	// 내부 이터레이터가 패턴의 마지막에 도달한 경우 리셋해야 한다.
	if p.i >= len(p.pattern) {
		p.i = 0
	}
	// 패턴과 문제의 답이 일치하는 경우 Count 를 증가시킨다.
	if answer == p.pattern[p.i] {
		p.count++
	}
	p.i++
}

func biggest(counts []int) []int {
	max := math.MinInt8
	biggest := make([]int, 0)

	// 가장 큰 값을 찾는다.
	for _, count := range counts {
		if max < count {
			max = count
		}
	}
	// 가장 큰 값이 몇개가 들어있는지 찾는다.
	for i, count := range counts {
		if max == count {
			biggest = append(biggest, i+1)
		}
	}
	return biggest
}

func solution(answers []int) []int {
	persons := []*p{
		{[]int{1,2,3,4,5}, 0, 0},
		{[]int{2,1,2,3,2,4,2,5}, 0, 0},
		{[]int{3,3,1,1,2,2,4,4,5,5}, 0, 0},
	}
	counts := make([]int, 0)

	for _, answer := range answers {
		for _, p := range persons {
			p.resolve(answer)
		}
	}

	for _, p := range persons {
		counts = append(counts, p.count)
	}

	return biggest(counts)
}
