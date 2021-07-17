package _76501

func solution(absolutes []int, signs []bool) int {
	var sum int

	for k, abs := range absolutes {
		if signs[k] {
			sum += abs
		} else {
			sum -= abs
		}
	}

	return sum
}
