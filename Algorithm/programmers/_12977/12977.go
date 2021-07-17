// Package _12977
// https://programmers.co.kr/learn/courses/30/lessons/12977
package _12977

func isPrime(n int) bool {
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func solution(nums []int) int {
	var answer int

	for f := range nums {
		for s := range nums {
			if f != s {
				for i := range nums {
					if i != f && i != s && isPrime(nums[f]+nums[s]+nums[i]) {
						answer++
					}
				}
			}
		}
	}

	return answer / 6
}
