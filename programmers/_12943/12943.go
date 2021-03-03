package _12943

func solution(num int) int {
	var answer int

	for num != 1 {
		if num % 2 == 0 {
			num /= 2
		} else {
			num = (num * 3) + 1
		}
		answer++
	}
	if answer > 500 {
		return -1
	}

	return answer
}