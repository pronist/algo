package _42840

import (
	"testing"
)

func TestBiggest(t *testing.T) {
	cases := []struct {
		counts []int
		expect []int
	}{
		{[]int{7, 0, 0}, []int{1}},
		{[]int{2, 2, 2}, []int{1, 2, 3}},
	}
	for _, c := range cases {
		biggest := Biggest(c.counts)
		if len(biggest) != len(c.expect) {
			t.Errorf("got %#v, want %#v", biggest, c.expect)
			return
		}
		for i, b := range biggest {
			if b != c.expect[i] {
				t.Errorf("got %#v, want %#v", biggest, c.expect)
			}
		}
	}
}

func TestP_Resolve(t *testing.T) {
	cases := []struct {
		P       P
		answers []int
		expect  int
	}{
		{P{[]int{1, 2, 3, 4, 5}, 0, 0}, []int{1, 2, 3, 4, 5}, 5},
		{P{[]int{2, 1, 2, 3, 2, 4, 2, 5}, 0, 0}, []int{1, 2, 3, 4, 5}, 0},
		{P{[]int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}, 0, 0}, []int{1, 2, 3, 4, 5}, 0},
		{P{[]int{1, 2, 3, 4, 5}, 0, 0}, []int{1, 3, 2, 4, 2}, 2},
		{P{[]int{2, 1, 2, 3, 2, 4, 2, 5}, 0, 0}, []int{1, 3, 2, 4, 2}, 2},
		{P{[]int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}, 0, 0}, []int{1, 3, 2, 4, 2}, 2},
	}
	for _, c := range cases {
		for _, answer := range c.answers {
			c.P.Resolve(answer)
		}
		if c.P.Count != c.expect {
			t.Errorf("got %#v, want %#v", c.P.Count, c.expect)
		}
	}
}

func TestSolution(t *testing.T) {
	cases := []struct {
		answers []int
		expect  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1}},
		{[]int{1, 3, 2, 4, 2}, []int{1, 2, 3}},
	}
	for _, c := range cases {
		biggest := solution(c.answers)
		if len(biggest) != len(c.expect) {
			t.Errorf("got %#v, want %#v", biggest, c.expect)
			return
		}
		for i, b := range biggest {
			if b != c.expect[i] {
				t.Errorf("got %#v, want %#v", biggest, c.expect)
			}
		}
	}
}
