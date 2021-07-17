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
		biggest := biggest(c.counts)
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
		p       p
		answers []int
		expect  int
	}{
		{p{[]int{1, 2, 3, 4, 5}, 0, 0}, []int{1, 2, 3, 4, 5}, 5},
		{p{[]int{2, 1, 2, 3, 2, 4, 2, 5}, 0, 0}, []int{1, 2, 3, 4, 5}, 0},
		{p{[]int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}, 0, 0}, []int{1, 2, 3, 4, 5}, 0},
		{p{[]int{1, 2, 3, 4, 5}, 0, 0}, []int{1, 3, 2, 4, 2}, 2},
		{p{[]int{2, 1, 2, 3, 2, 4, 2, 5}, 0, 0}, []int{1, 3, 2, 4, 2}, 2},
		{p{[]int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}, 0, 0}, []int{1, 3, 2, 4, 2}, 2},
	}
	for _, c := range cases {
		for _, answer := range c.answers {
			c.p.resolve(answer)
		}
		if c.p.count != c.expect {
			t.Errorf("got %#v, want %#v", c.p.count, c.expect)
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
