// Package _42840
// https://programmers.co.kr/learn/courses/30/lessons/42840
package _42840

import (
	"math"
)

type P struct {
	pattern []int
	i       int
	Count   int
}

func (p *P) Resolve(answer int) {
	if p.i >= len(p.pattern) {
		p.i = 0
	}
	if answer == p.pattern[p.i] {
		p.Count++
	}
	p.i++
}

func biggest(counts []int) []int {
	max := math.MinInt8
	biggest := make([]int, 0)

	for _, count := range counts {
		if max < count {
			max = count
		}
	}
	for i, count := range counts {
		if max == count {
			biggest = append(biggest, i+1)
		}
	}
	return biggest
}

func solution(answers []int) []int {
	persons := []*P{
		{[]int{1, 2, 3, 4, 5}, 0, 0},
		{[]int{2, 1, 2, 3, 2, 4, 2, 5}, 0, 0},
		{[]int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}, 0, 0},
	}
	for _, answer := range answers {
		for _, p := range persons {
			p.Resolve(answer)
		}
	}

	return biggest([]int{persons[0].Count, persons[1].Count, persons[2].Count})
}
