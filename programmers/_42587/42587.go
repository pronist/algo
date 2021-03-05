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
				// 큐의 맨 끝으로 보낸다.
				queue = append(queue, queue[0])
				queue = queue[1:]

				// 규칙에 의한 target 처리
				switch target {
				case 0:
					target = len(queue) - 1
				default:
					target--
				}

				// 이미 처리를 했다면 더 이상 진행할 필요가 없으므로 아래의 코드는 생략.
				continue Loop
			}
		}
		// 여기서 target 이 0 이 된다는 것은 큐에서 target 을 제외하고 모두 빠져나갔다는 의미다.
		if target == 0 {
			break Loop
		}
		// 큐에서 원소를 하나 제거하고 target, answer 처리.
		queue = queue[1:]
		target, answer = target-1, answer+1
	}

	return answer + 1
}
