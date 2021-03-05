// Package _42587
// https://programmers.co.kr/learn/courses/30/lessons/42587
package _42587

func solution(priorities []int, location int) int {
	var answer int

	queue := make([]int, len(priorities))
	copy(queue, priorities)

	target := location

Loop:
	for {
		for _, priority := range queue {
			if priority > queue[0] {
				queue = append(queue, queue[0])
				queue = queue[1:]

				switch target {
				case 0:
					target = len(queue) - 1
				default:
					target--
				}
				continue Loop
			}
		}
		if target == 0 {
			break Loop
		}
		queue = queue[1:]
		target, answer = target-1, answer+1
	}

	return answer + 1
}
