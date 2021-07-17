// Package _12901
// https://programmers.co.kr/learn/courses/30/lessons/12901
package _12901

//import (
//	"strings"
//	"time"
//)

var weekday = []string{
	"SUN",
	"MON",
	"TUE",
	"WED",
	"THU",
	"FRI",
	"SAT",
}

var DayOfMonth = []int{
	31,
	29, // 2016년은 윤년이다.
	31,
	30,
	31,
	30,
	31,
	31,
	30,
	31,
	30,
	31,
}

func solution(a int, b int) string {
	/*
	 * Solution 1
	 */
	//t := time.Date(2016, time.Month(a), b, 0, 0, 0, 0, time.UTC)
	//return strings.ToUpper(t.Weekday().String()[:3])

	/*
	 * Solution 2
	 */
	var d, w int

	for i := 0; i < a-1; i++ {
		d += DayOfMonth[i]
	}
	// 현재 주어진 날짜 이전까지 더한다.
	d += b-1
	// 5의 의미는 weekday 슬라이스에서 'FRI' 의 값이 위치한다.
	// 2016년 1월 1일은 금요일이기 때문에 더해야 한다.
	w = d%7 + 5
	// w 는 weekday 의 인덱스를 의미하기도 하기 때문에 넘어가지 않도록 한다.
	if w >= len(weekday) {
		w -= 7
	}

	return weekday[w]
}
