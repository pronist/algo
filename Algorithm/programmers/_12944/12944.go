package _12944

func solution(arr []int) float64 {
	var answer float64

	// reduce
	for _, v := range arr {
		answer += float64(v)
	}

	return answer / float64(len(arr))
}