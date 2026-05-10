package main

import "fmt"

// 수학, 영어, 역사 점수를 받아서 합계(total)와 평균(avg)을 반환합니다.
func getGrade(math, eng, history int) (total int, avg float64) {
	total = math + eng + history
	avg = float64(total) / 3.0
	return
}

func main() {
	// 1. 과목별 점수 설정
	mathScore := 85
	engScore := 92
	historyScore := 78

	// 2. 함수 호출: 두 개의 결과값을 한 번에 받음
	total, avg := getGrade(mathScore, engScore, historyScore)

	// 3. 결과 출력
	fmt.Println("--- 성적 결과 리포트 ---")
	fmt.Printf("수학: %d점, 영어: %d점, 역사: %d점\n", mathScore, engScore, historyScore)
	fmt.Printf("총점: %d점\n", total)
	fmt.Printf("평균: %.2f점\n", avg)
}
